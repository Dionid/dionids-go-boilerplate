package http

import (
	"context"

	"github.com/Dionid/go-boiler/api/v1/go/proto"
	"github.com/Dionid/go-boiler/features"
	fsignin "github.com/Dionid/go-boiler/features/sign-in"
	fsignup "github.com/Dionid/go-boiler/features/sign-up"
)

type MainApiService struct {
	proto.UnimplementedMainApiServer
	Deps *features.Deps
}

// # Auth

func (service *MainApiService) SignIn(ctx context.Context, request *proto.SignInCallRequest) (*proto.SignInCallResponse, error) {
	return fsignin.SignIn(ctx, service.Deps, request)
}

func (service *MainApiService) SignUp(ctx context.Context, request *proto.SignUpCallRequest) (*proto.SignUpCallResponse, error) {
	return fsignup.SignUp(ctx, service.Deps, request)
}
