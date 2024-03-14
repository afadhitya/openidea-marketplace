package signin

import (
	"context"
	"errors"
	"log"

	userrepositories "github.com/widcha/openidea-marketplace/internal/app/modules/user/repositories"
	"github.com/widcha/openidea-marketplace/internal/pkg/pwd"
	"github.com/widcha/openidea-marketplace/internal/pkg/token"
)

type interactor struct {
	userRepo userrepositories.IRepo
	jwtToken token.JwtCreateToken
}

func NewUsecase(userRepo userrepositories.IRepo, jwtToken token.JwtCreateToken) Inport {
	return &interactor{
		userRepo: userRepo,
		jwtToken: jwtToken,
	}
}

func (i interactor) Execute(ctx context.Context, req InportRequest) (InportResponse, error) {
	user, err := i.userRepo.User().GetbyUsername(ctx, req.Username)
	if err != nil {
		log.Println(err)
		return InportResponse{}, err
	}

	passTrue := pwd.ComparePasswords(user.Password, []byte(req.Password))
	if !passTrue {
		log.Println(err)
		return InportResponse{}, errors.New("password incorrect")
	}

	userToken, _, err := i.jwtToken.CreateTokenUser(ctx, user)
	if err != nil {
		log.Println(err)
		return InportResponse{}, err
	}

	return InportResponse{
		Token:    userToken,
		Username: user.Username,
		Name:     user.Name,
	}, nil
}
