package jwttoken

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kooroshh/fiber-boostrap/pkg/env"
)

type ClaimsToken struct {
	Username string `json:"username"`
	Fullname string `json:"full_name"`
	jwt.RegisteredClaims
}

var MapToken = map[string]time.Duration{
	"token":         3 * time.Hour,
	"refresh_token": 72 * time.Hour,
}

var secretKey = []byte(env.GetEnv("APP_SECRET", ""))

func GeneratedToken(ctx context.Context, username, fullname string, tokenType string) (string, error) {
	claimsToken := ClaimsToken{
		Username: username,
		Fullname: fullname,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    env.GetEnv("APP_NAME", ""),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(MapToken[tokenType])),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsToken)
	result, err := token.SignedString(secretKey)
	if err != nil {
		return result, fmt.Errorf("failed generated token %s", err)
	}
	return result, nil
}

func ValidateToken(ctx context.Context, token string) (*ClaimsToken, error) {
	var (
		claimToken *ClaimsToken
		ok         bool
	)

	jwtToken, err := jwt.ParseWithClaims(token, &ClaimsToken{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed validate method jwt : %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt: %v", err)
	}

	if claimToken, ok = jwtToken.Claims.(*ClaimsToken); !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("token invalid : %v", err)
	}

	return claimToken, nil
}
