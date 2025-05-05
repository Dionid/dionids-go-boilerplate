package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	_ "github.com/bufbuild/protovalidate-go"
	_ "github.com/lib/pq"

	"github.com/Dionid/go-boiler/features"
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

	// # Graceful shutdown emitter
	gse := make(chan string, 1)

	// ## Sub to SIGTERM and SIGINT
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigs
		gse <- sig.String()
	}()

	// # Global WaitGroup
	gwg := &sync.WaitGroup{}

	// # Global Context
	ctx, cancel := context.WithCancel(context.Background())

	// # DB
	mainPgPool, err := sqlx.Open("postgres", config.MainDbConnectionStr)
	if err != nil {
		log.Fatalf("DB error: %v\n", err)
	}
	defer mainPgPool.Close()

	mainPgPool.SetMaxOpenConns(10)

	// # Init first admin
	err = initFirstAdmin(ctx, config, mainPgPool)
	if err != nil {
		if v, ok := err.(*terrors.PrivateError); ok {
			log.Fatalf("Init first admin: %v\n", v.GetPrivateMessage())
		}
		log.Fatalf("Init first admin: %v\n", err)
	}

	// # Deps
	deps := &features.Deps{
		Logger: logger,
		MainDb: mainPgPool,
		Config: features.Config{
			JwtSecret:       []byte(config.JwtSecret),
			ExpireInSeconds: config.JwtExpireInSeconds,
		},
		GlobalWg:                gwg,
		GracefulShutdownEmitter: gse,
	}

	// # Server
	e, serveGrpc, serveCmux, closeServer, err := initServer(ctx, config, logger, deps)
	if err != nil {
		log.Fatal(err)
	}

	gwg.Add(1)
	go func() {
		logger.Info(fmt.Sprintf("Starting gRPC Gateway on port %d", config.Port))
		err := e.Start(fmt.Sprintf(":%d", config.Port))
		logger.Info(fmt.Sprintf("echo %v", err))
		if err != nil && !strings.Contains(err.Error(), "server closed") && err != http.ErrServerClosed {
			logger.Sugar().Errorf("e.Start err: %v\n", err)
			log.Fatal(err)
		}

		gwg.Done()
	}()

	gwg.Add(1)
	go func() {
		logger.Info(fmt.Sprintf("Starting gRPC Node on port %d", config.Port))
		err := serveGrpc()
		logger.Info(fmt.Sprintf("grpc %v", err))
		if err != nil && !strings.Contains(err.Error(), "server closed") {
			logger.Sugar().Errorf("serveGrpc err: %v\n", err)
			log.Fatal(err)
		}

		gwg.Done()
	}()

	gwg.Add(1)
	go func() {
		logger.Info(fmt.Sprintf("Starting combined server on port %d", config.Port))
		err := serveCmux()
		logger.Info(fmt.Sprintf("cmux: %v", err))
		if err != nil && !strings.Contains(err.Error(), "closed network connection") {
			logger.Sugar().Errorf("serveCmux err: %v\n", err)
			log.Fatal(err)
		}

		gwg.Done()
	}()

	// # Graceful shutdown
	go func() {
		reason := <-gse
		gwg.Add(1)
		logger.Info(fmt.Sprintf("Received %s signal", reason))
		closeServer()
		logger.Info("Server closed")
		mainPgPool.Close()
		logger.Info("mainPgPool closed")
		cancel()
		logger.Info("ctx canceled")
		gwg.Done()
	}()

	logger.Info("Started")
	gwg.Wait()
	logger.Info("Bye")
}
