package app

import (
	"github.com/widcha/openidea-marketplace/internal/app/modules/health"
	"github.com/widcha/openidea-marketplace/internal/pkg"
)

type Container struct {
	HealthCheckUsecase health.Usecase
}

func NewContainer(datasource *pkg.Datasource) *Container {
	return &Container{
		HealthCheckUsecase: health.NewUsecase(datasource.Postgre),
	}
}
