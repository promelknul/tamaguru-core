package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var signingKey = []byte("replace-me")

type Claims struct {
	UserID string `json:"uid"`
	jwt.RegisteredClaims
}

func Generate(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signingKey)
}
