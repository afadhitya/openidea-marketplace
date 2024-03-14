package token

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/widcha/openidea-marketplace/configs"
	userentities "github.com/widcha/openidea-marketplace/internal/app/modules/user"
)

type JwtCreateToken interface {
	CreateTokenUser(ctx context.Context, user userentities.User) (string, int64, error)
}

type JWTToken struct {
	secretKey string
}

func NewJWTToken(secretKey string) *JWTToken {
	if strings.TrimSpace(secretKey) == "" {
		return nil
	}

	return &JWTToken{
		secretKey: secretKey,
	}
}

func (g *JWTToken) CreateTokenUser(ctx context.Context, user userentities.User) (string, int64, error) {
	expiredTime := time.Now().Add(time.Second * time.Duration(configs.Get().MaxAgeToken)).Unix()

	claims := jwt.MapClaims{
		"userid":   user.Id,
		"name":     user.Name,
		"username": user.Username,
		"exp":      expiredTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configs.Get().JwtSecret))
	if err != nil {
		log.Println(err)
		return "", expiredTime, err
	}

	return tokenString, expiredTime, nil
}
