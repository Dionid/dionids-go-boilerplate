package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Dionid/go-boiler/api/v1/go/proto"
	_ "github.com/bufbuild/protovalidate-go"
	_ "github.com/lib/pq"

	"github.com/Dionid/go-boiler/dbs/maindb"

	"github.com/Dionid/go-boiler/features"
	fsignin "github.com/Dionid/go-boiler/features/sign-in"
	"github.com/Dionid/go-boiler/pkg/terrors"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var allowedHeaders = map[string]struct{}{
	"x-request-id": {},
}

func isHeaderAllowed(s string) (string, bool) {
	// check if allowedHeaders contain the header
	if _, isAllowed := allowedHeaders[s]; isAllowed {
		// send uppercase header
		return strings.ToUpper(s), true
	}
	// if not in the allowed header, don't send the header
	return s, false
}

func main() {
	// # Config
	config, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	// # Logger
	zapConfig := zap.NewDevelopmentConfig()
	if config.Env == "production" {
		zapConfig = zap.NewProductionConfig()
	}
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Starting")

	// # Graceful shutdown sigs
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	// # Global WaitGroup
	wg := sync.WaitGroup{}

	// # Global Context
	ctx, cancel := context.WithCancel(context.Background())

	// # DB
	mainPgPool, err := sqlx.Open("postgres", config.MainDbConnectionStr)
	if err != nil {
		log.Fatalf("DB error: %v\n", err)
	}
	defer mainPgPool.Close()

	mainPgPool.SetMaxOpenConns(10)

	mainDbQueries := maindb.New(mainPgPool)

	// # Init first admin
	err = initFirstAdmin(ctx, config, mainPgPool)
	if err != nil {
		if v, ok := err.(*terrors.PrivateError); ok {
			log.Fatalf("Init first admin: %v\n", v.GetPrivateMessage())
		}
		log.Fatalf("Init first admin: %v\n", err)
	}

	transport, err := initRmq(ctx, config)
	if err != nil {
		log.Fatalf("Transport error: %v\n", err)
	}

	logger.Info("RMQ")

	// # Deps
	deps := &features.Deps{
		Logger:        logger,
		MainDb:        mainPgPool,
		MainDbQueries: mainDbQueries,
		Config: features.Config{
			JwtSecret:       []byte(config.JwtSecret),
			ExpireInSeconds: config.JwtExpireInSeconds,
		},
		RmqT: transport,
	}

	err = transport.SubscribeOnCall(
		ctx,
		&proto.SignInCallRequest{
			Name: "SignIn",
		},
		func(ctx context.Context, requestBody []byte) ([]byte, terrors.Error) {
			request := &proto.SignInCallRequest{}
			jErr := json.Unmarshal(requestBody, request)
			if jErr != nil {
				return nil, terrors.NewPrivateError("failed to unmarshal request")
			}

			result, err := fsignin.SignIn(ctx, deps, request)
			if err != nil {
				return nil, err
			}

			response, jErr := json.Marshal(result)
			if jErr != nil {
				return nil, terrors.NewPrivateError("failed to marshal response")
			}

			if err != nil {
				return nil, terrors.NewPrivateError("failed to marshal response")
			}

			return response, nil
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	// # Server
	e, serveGrpc, serveCmux, closeServer, err := initServer(ctx, config, logger, deps)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func() {
		logger.Info(fmt.Sprintf("Starting gRPC Gateway on port %d", config.Port))
		err := e.Start(fmt.Sprintf(":%d", config.Port))
		logger.Info(fmt.Sprintf("echo %v", err))
		if err != nil && !strings.Contains(err.Error(), "server closed") && err != http.ErrServerClosed {
			logger.Sugar().Errorf("e.Start err: %v\n", err)
			log.Fatal(err)
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		logger.Info(fmt.Sprintf("Starting gRPC Node on port %d", config.Port))
		err := serveGrpc()
		logger.Info(fmt.Sprintf("grpc %v", err))
		if err != nil && !strings.Contains(err.Error(), "server closed") {
			logger.Sugar().Errorf("serveGrpc err: %v\n", err)
			log.Fatal(err)
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		logger.Info(fmt.Sprintf("Starting combined server on port %d", config.Port))
		err := serveCmux()
		logger.Info(fmt.Sprintf("cmux: %v", err))
		if err != nil && !strings.Contains(err.Error(), "closed network connection") {
			logger.Sugar().Errorf("serveCmux err: %v\n", err)
			log.Fatal(err)
		}

		wg.Done()
	}()

	go func() {
		sig := <-sigs
		wg.Add(1)
		logger.Info(fmt.Sprintf("Received %s signal", sig))
		closeServer()
		logger.Info("Server closed")
		mainPgPool.Close()
		logger.Info("mainPgPool closed")
		transport.Close()
		logger.Info("transport closed")
		cancel()
		logger.Info("ctx canceled")
		wg.Done()
	}()

	logger.Info("Started")
	wg.Wait()
	logger.Info("Bye")
}
