package userusecase

import (
	"context"
	"log"

	"github.com/google/uuid"
	userentities "github.com/widcha/openidea-marketplace/internal/app/modules/user"
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
	hashPass := pwd.HashAndSalt([]byte(req.Password))

	user := userentities.User{
		Id:       uuid.NewString(),
		Name:     req.Name,
		Username: req.Username,
		Password: hashPass,
	}

	err := i.userRepo.User().Create(ctx, user)
	if err != nil {
		log.Println(err)
		return InportResponse{}, err
	}

	userToken, _, err := i.jwtToken.CreateTokenUser(ctx, user)
	if err != nil {
		log.Fatalln(err)
		return InportResponse{}, err
	}

	return InportResponse{
		Token:    userToken,
		Username: user.Username,
		Name:     user.Name,
	}, nil
}
