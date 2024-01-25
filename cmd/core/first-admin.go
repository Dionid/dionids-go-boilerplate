package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/Dionid/go-boiler/dbs/maindb"
	"github.com/Dionid/go-boiler/pkg/ff"
	"github.com/Dionid/go-boiler/pkg/terrors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func initFirstAdmin(ctx context.Context, config *Config, db *sqlx.DB) terrors.Error {
	userFound, err := maindb.SelectUserTableByEmail(ctx, db, config.FirstUserEmail)
	if err != nil {
		if !terrors.IsNotFoundErr(err) {
			return terrors.NewDbErr(err)
		}
	}

	if userFound != nil && userFound.Email == config.FirstUserEmail {
		return nil
	}

	email := config.FirstUserEmail
	password := config.FirstUserPassword

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return terrors.NewPrivateError("can't generate password hash")
	}

	newUser := maindb.NewInsertableUserModel(
		uuid.New(),
		email,
		string(hashedPassword),
		time.Now(),
		sql.NullTime{},
		"admin",
	)

	result, err := maindb.InsertIntoUserTable(ctx, db, newUser)
	if err != nil {
		return terrors.NewDbErr(err)
	}

	if ff.IgnoreError(result.RowsAffected()) != 1 {
		return terrors.NewPrivateError("can't insert user")
	}

	return nil
}
