package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"text/template"

	"github.com/Dionid/go-boiler/api/v1/go/proto"
	httpapi "github.com/Dionid/go-boiler/cmd/core/http"
	"github.com/Dionid/go-boiler/features"
	"github.com/Dionid/go-boiler/pkg/terrors"
	"github.com/brpaz/echozap"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func mapError(err error, logger *zap.Logger) terrors.Error {
	switch v := err.(type) {
	case terrors.Error:
		logger.Error(v.GetPrivateMessage())
		return v
	default:
		return terrors.NewPrivateError(v.Error())
	}
}

func initServer(ctx context.Context, config *Config, logger *zap.Logger, deps *features.Deps) (e *echo.Echo, serveGRPC func() error, serveHTTP func() error, Close func(), err error) {
	e = echo.New()

	pprof.Register(e)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(echozap.ZapLogger(logger))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		mapedErr := mapError(err, logger)
		c.JSON(mapedErr.GetCode(), map[string]interface{}{
			"message": mapedErr.Error(),
			"code":    mapedErr.GetCode(),
		})
	}

	// # gRPC
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(InterceptorLogger(logger), logging.WithLogOnEvents(logging.StartCall, logging.FinishCall)),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(InterceptorLogger(logger), logging.WithLogOnEvents(logging.StartCall, logging.FinishCall)),
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	)
	proto.RegisterMainApiServer(grpcServer, &httpapi.MainApiService{Deps: deps})

	// # gRPC Gateway
	mux := runtime.NewServeMux(
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			// send all the headers received from the client
			md := metadata.Pairs("auth", header)
			return md
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			mapedErr := mapError(err, logger)

			fmt.Printf("err: %+v\n", err)

			//creating a new HTTTPStatusError with a custom status, and passing error
			newError := runtime.HTTPStatusError{
				HTTPStatus: mapedErr.GetCode(),
				Err:        mapedErr,
			}

			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &newError)
		}),
	)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err = proto.RegisterMainApiHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%d", config.Host, config.Port), opts)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Creating a normal HTTP server
	e.GET("/", httpapi.HealthCheck)

	// # Swagger
	e.GET("/swagger", func(req echo.Context) error {
		path := fmt.Sprintf("%s%s", config.SwaggerPathPrefix, "/swagger.html")

		t := template.Must(template.ParseGlob(path))

		var result bytes.Buffer

		httpType := "http"
		if config.Https {
			httpType = "https"
		}

		err := t.Execute(&result, map[string]interface{}{
			"Url": fmt.Sprintf("%s://%s:%d%s", httpType, config.Host, config.Port, "/openapi/v1/openapi.yaml"),
		})

		if err != nil {
			return err
		}

		return req.HTML(http.StatusOK, result.String())
	})
	e.File("/openapi/v1/openapi.yaml", fmt.Sprintf("%s%s", config.SwaggerPathPrefix, "/openapi.yaml"))

	// # gRPC Gateway
	e.Group("/api/v1/*{grpc_gateway}").Any("", echo.WrapHandler(mux))

	// creating a listener for server
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	m := cmux.New(l)

	// a different listener for HTTP1
	httpL := m.Match(cmux.HTTP1Fast())

	// a different listener for HTTP2 since gRPC uses HTTP2
	grpcL := m.Match(cmux.HTTP2())

	// start server
	// passing dummy listener
	e.Listener = httpL

	return e,
		func() error { return grpcServer.Serve(grpcL) },
		func() error { return m.Serve() },
		func() {
			m.Close()
			e.Shutdown(ctx)
			grpcL.Close()
		},
		nil
}
