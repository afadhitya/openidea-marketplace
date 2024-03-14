package app

import (
	"github.com/widcha/openidea-marketplace/configs"
	"github.com/widcha/openidea-marketplace/internal/app/modules/health"
	userrepositories "github.com/widcha/openidea-marketplace/internal/app/modules/user/repositories"
	"github.com/widcha/openidea-marketplace/internal/app/modules/user/usercase/signin"
	sign "github.com/widcha/openidea-marketplace/internal/app/modules/user/usercase/signin"
	userusecase "github.com/widcha/openidea-marketplace/internal/app/modules/user/usercase/signup"
	"github.com/widcha/openidea-marketplace/internal/pkg"
	"github.com/widcha/openidea-marketplace/internal/pkg/token"
)

type Container struct {
	HealthCheckUsecase health.Usecase
	UserSignupUsecase  userusecase.Inport
	UserSigninUsecase  sign.Inport
}

func NewContainer(datasource *pkg.Datasource) *Container {
	userRepo := userrepositories.NewRepo(datasource)
	jwtlib := token.NewJWTToken(configs.Get().JwtSecret)

	return &Container{
		HealthCheckUsecase: health.NewUsecase(datasource.Postgre),
		UserSignupUsecase:  userusecase.NewUsecase(userRepo, jwtlib),
		UserSigninUsecase:  signin.NewUsecase(userRepo, jwtlib),
	}
}
