package inttests

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Dionid/go-boiler/features"
	"github.com/Dionid/go-boiler/pkg/df"
	"go.uber.org/zap"
)

type TestDeps struct {
	Config           *TestConfig
	Logger           *zap.Logger
	MainDbConnection *sqlx.DB
	FeaturesConfig   features.Config
	RmqTransport     *df.RmqTransport
	Cleanup          func() error
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func createTemplateDb(
	ctx context.Context,
	connectionStr string,
	mainDbName string,
	tempDbName string,
	attempt int,
) error {
	if attempt > 10 {
		return fmt.Errorf("createTemplateDb: attempt limit exceeded")
	}

	conn, err := sqlx.Open("postgres", connectionStr)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(ctx, fmt.Sprintf(`CREATE DATABASE "%s" TEMPLATE "%s";`, tempDbName, mainDbName)); err != nil {
		if strings.Contains(err.Error(), "is being accessed by other users") {
			conn.Close()
			dur := time.Duration(rand.Intn(1000)) * time.Millisecond
			time.Sleep(dur)
			return createTemplateDb(ctx, connectionStr, mainDbName, tempDbName, attempt+1)
		}
		return err
	}
	return nil
}

func dropTemplateTable(
	ctx context.Context,
	connectionStr string,
	tempDbName string,
) error {
	conn, err := sqlx.Open("postgres", connectionStr)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(ctx, fmt.Sprintf(`DROP DATABASE "%s";`, tempDbName)); err != nil {
		return err
	}

	return nil
}

func InitTestDeps(ctx context.Context) (*TestDeps, error) {
	config, err := NewTestConfig()
	if err != nil {
		return nil, err
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	tempDbName := "go-boiler-" + RandStringRunes(10)

	if err = createTemplateDb(ctx, config.MainDbConnection, "go-boiler", tempDbName, 0); err != nil {
		return nil, err
	}

	mainDbConnectionTemplate, err := sqlx.Open("postgres", strings.Replace(config.MainDbConnection, "go-boiler", tempDbName, 1))
	if err != nil {
		return nil, err
	}

	rmqTransport, err := df.NewRmqTransport(
		func(transport df.RmqTransport) (df.RmqTransport, error) {
			transport.ConnectionString = config.RmqConnection

			return transport, nil
		},
	)

	featuresConfig := features.Config{
		JwtSecret:       []byte("secret"),
		ExpireInSeconds: 10000,
	}

	result := &TestDeps{
		config,
		logger,
		mainDbConnectionTemplate,
		featuresConfig,
		rmqTransport,
		func() error {
			mainDbConnectionTemplate.Close()
			if err = dropTemplateTable(ctx, config.MainDbConnection, tempDbName); err != nil {
				return err
			}
			return nil
		},
	}

	return result, nil
}
