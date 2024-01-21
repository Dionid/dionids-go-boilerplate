package features

import (
	"github.com/Dionid/go-boiler/dbs/maindb"
	"github.com/Dionid/go-boiler/pkg/df"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Config struct {
	JwtSecret       []byte
	ExpireInSeconds int64
}

type Deps struct {
	Logger *zap.Logger

	MainDb        *sqlx.DB
	MainDbQueries *maindb.Queries

	Config Config

	RmqT *df.RmqTransport
}
