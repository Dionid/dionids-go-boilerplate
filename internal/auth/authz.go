package auth

import (
	"github.com/Dionid/go-boiler/api/v1/go/proto"
	"github.com/Dionid/go-boiler/pkg/terrors"
)

type Request interface {
	GetMeta() *proto.Meta
}

func AuthorizeByRoles(jwtSecret []byte, roles []string, request Request) error {
	meta := request.GetMeta()

	if meta == nil {
		return terrors.NewUnauthorizedError("meta is required", nil)
	}

	if meta.Token == nil {
		return terrors.NewUnauthorizedError("token is required", nil)
	}

	claims, err := ParseToken(jwtSecret, *meta.Token)
	if err != nil {
		return terrors.NewUnauthorizedError("invalid token", nil)
	}

	for _, role := range roles {
		if role == claims.Role {
			return nil
		}
	}

	return terrors.NewForbiddenError("invalid token", nil)
}
