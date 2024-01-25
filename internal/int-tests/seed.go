package inttests

import (
	"context"
	"database/sql"
	"time"

	"github.com/Dionid/go-boiler/dbs/maindb"
	"github.com/Dionid/go-boiler/features"
	"github.com/Dionid/go-boiler/internal/auth"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type SeedResult struct {
	User     *maindb.UserModel
	JwtToken string
}

func Seed(
	ctx context.Context,
	featureConfig features.Config,
	mainConn *sqlx.DB,
) (*SeedResult, error) {
	email := "dio@mail.com"
	password := "1234"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil
	}

	userModel := maindb.NewInsertableUserModel(
		uuid.New(),
		"Admin",
		email,
		string(hashedPassword),
		time.Now(),
		sql.NullTime{},
		"admin",
	)

	user, err := maindb.InsertIntoUserTableReturningAll(
		ctx,
		mainConn,
		userModel,
	)
	if err != nil {
		return nil, err
	}

	jwtToken, err := auth.CreateToken(featureConfig.JwtSecret, featureConfig.ExpireInSeconds, user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	result := &SeedResult{
		user,
		jwtToken,
	}

	return result, nil
}
