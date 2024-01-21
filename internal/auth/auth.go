package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Claims struct {
	UserId    uuid.UUID `json:"sid"`
	Role      string    `json:"role"`
	ExpiresAt int64     `json:"exp,omitempty"`
}

func verifyExp(exp int64, now int64, required bool) bool {
	if exp == 0 {
		return !required
	}
	return now <= exp
}

func (c *Claims) Valid() error {
	vErr := new(jwt.ValidationError)
	now := time.Now().Unix()

	// The claims below are optional, by default, so if they are set to the
	// default value in Go, let's not fail the verification for them.
	if !verifyExp(c.ExpiresAt, now, false) {
		delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		vErr.Inner = fmt.Errorf("token is expired by %v", delta)
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	return nil
}

func CreateToken(jwtSecret []byte, expireInSeconds int64, userId uuid.UUID, userRole string) (string, error) {
	claims := &Claims{
		UserId:    userId,
		Role:      userRole,
		ExpiresAt: time.Now().Add(time.Duration(expireInSeconds) * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

func ParseToken(jwtSecret []byte, tokenString string) (*Claims, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}
