package features

import (
	"sync"

	"github.com/Dionid/go-boiler/pkg/df"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Config struct {
	JwtSecret       []byte
	ExpireInSeconds int64
}

type Deps struct {
	Logger                  *zap.Logger
	GlobalWg                *sync.WaitGroup
	GracefulShutdownEmitter chan string

	MainDb *sqlx.DB

	Config Config

	RmqT *df.RmqTransport
}
