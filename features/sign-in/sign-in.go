package fsignin

import (
	"context"

	"github.com/Dionid/go-boiler/api/v1/go/proto"
	"github.com/Dionid/go-boiler/features"
	"github.com/Dionid/go-boiler/internal/auth"
	"github.com/Dionid/go-boiler/pkg/terrors"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(ctx context.Context, deps *features.Deps, request *proto.SignInCallRequest) (*proto.SignInCallResponse, terrors.Error) {
	// # Validate request
	if request.Params.Email == "" {
		return nil, terrors.NewValidationError("NewValidationError", nil)
	}

	if request.Params.Password == "" {
		return nil, terrors.NewValidationError("NewValidationError", nil)
	}

	// # Query user
	user, err := deps.MainDbQueries.SignInGetUser(ctx, request.Params.Email)
	if err != nil {
		return nil, terrors.NewValidationError("Incorrect email or password", nil)
	}

	// # Hash and compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Params.Password))
	if err != nil {
		return nil, terrors.NewValidationError("Incorrect email or password", nil)
	}

	tokenString, err := auth.CreateToken(deps.Config.JwtSecret, deps.Config.ExpireInSeconds, user.ID, user.Role)
	if err != nil {
		return nil, terrors.NewPrivateError("in create token")
	}

	// TODO: # Add session to redis
	// ...

	resp := &proto.SignInCallResponse{
		Id: request.Id,
		Result: &proto.SignInCallResponse_Result{
			Result: &proto.SignInCallResponse_Result_Success_{
				Success: &proto.SignInCallResponse_Result_Success{
					Token: tokenString,
				},
			},
		},
	}

	return resp, nil
}
