package http

import (
	"context"

	"github.com/Dionid/go-boiler/api/v1/go/proto"
	"github.com/Dionid/go-boiler/features"
)

type MainApiService struct {
	proto.UnimplementedMainApiServer
	Deps *features.Deps
}

// # Auth

func (service *MainApiService) SignIn(ctx context.Context, request *proto.SignInCallRequest) (*proto.SignInCallResponse, error) {
	return features.SignIn(ctx, service.Deps, request)
}
