package auth_test

import (
	"testing"

	"github.com/Dionid/go-boiler/internal/auth"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUnitAuth(t *testing.T) {
	userId := uuid.New()
	token, err := auth.CreateToken([]byte("secret"), 10000, userId, "admin")
	assert.Nil(t, err)
	assert.NotNil(t, token)
	assert.Greater(t, len(token), 0)

	claims, err := auth.ParseToken([]byte("secret"), token)
	assert.Nil(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, userId.String(), claims.UserId.String())
	assert.Equal(t, "admin", claims.Role)
}
