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
	CreateTokenUser(ctx context.Context, user userentities.User) (string, time.Time, error)
}

type UserClaims struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.RegisteredClaims
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

func (g *JWTToken) CreateTokenUser(ctx context.Context, user userentities.User) (string, time.Time, error) {
	issuedAt := time.Now()
	expiredTime := issuedAt.Add(time.Second * time.Duration(configs.Get().MaxAgeToken))

	userClaims := UserClaims{
		UserID:   user.Id,
		Name:     user.Name,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	tokenString, err := token.SignedString([]byte(configs.Get().JwtSecret))
	if err != nil {
		log.Println(err)
		return "", expiredTime, err
	}

	return tokenString, expiredTime, nil
}
