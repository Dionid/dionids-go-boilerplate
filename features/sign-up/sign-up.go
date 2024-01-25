package fsignup

import (
	"context"
	"database/sql"
	"time"

	"github.com/Dionid/go-boiler/api/v1/go/proto"
	"github.com/Dionid/go-boiler/dbs/maindb"
	"github.com/Dionid/go-boiler/features"
	"github.com/Dionid/go-boiler/internal/auth"
	"github.com/Dionid/go-boiler/pkg/terrors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx context.Context, deps *features.Deps, request *proto.SignUpCallRequest) (*proto.SignUpCallResponse, terrors.Error) {
	// # Validate request
	if request.Params.Email == "" {
		return nil, terrors.NewValidationError("NewValidationError", nil)
	}

	if request.Params.Password == "" {
		return nil, terrors.NewValidationError("NewValidationError", nil)
	}

	// # Query user
	userExists, _ := maindb.SelectUserTableByEmail(ctx, deps.MainDb, request.Params.Email)
	if userExists != nil {
		return nil, terrors.NewValidationError("Incorrect email or password", nil)
	}

	// # Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, terrors.NewPrivateError(err.Error())
	}

	// # Create User
	newUser := maindb.NewInsertableUserModel(
		uuid.New(),
		request.Params.Email,
		string(hashedPassword),
		time.Time{},
		sql.NullTime{},
		"client",
	)

	if _, err := maindb.InsertIntoUserTable(ctx, deps.MainDb, newUser); err != nil {
		return nil, terrors.NewPrivateError("in insert user")
	}

	tokenString, err := auth.CreateToken(deps.Config.JwtSecret, deps.Config.ExpireInSeconds, newUser.ID, newUser.Role)
	if err != nil {
		return nil, terrors.NewPrivateError("in create token")
	}

	// TODO: # Add session to redis
	// ...

	resp := &proto.SignUpCallResponse{
		Id: request.Id,
		Result: &proto.SignUpCallResponse_Result{
			Result: &proto.SignUpCallResponse_Result_Success_{
				Success: &proto.SignUpCallResponse_Result_Success{
					Token: tokenString,
				},
			},
		},
	}

	return resp, nil
}
