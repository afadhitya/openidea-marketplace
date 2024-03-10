package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/widcha/openidea-marketplace/configs"
	"github.com/widcha/openidea-marketplace/internal/app"
	"github.com/widcha/openidea-marketplace/internal/app/router"
	"github.com/widcha/openidea-marketplace/internal/middleware"
	"github.com/widcha/openidea-marketplace/internal/pkg"
)

func init() {
	middleware.PrometheusRegisterInit()
}

func main() {
	configs.Load()
	ginRouter := gin.Default()
	datasource := pkg.NewDataSource()
	container := app.NewContainer(datasource)
	router := router.NewRouter(ginRouter, datasource, container)
	router.RegisterRouter()

	ginRouter.Run(":" + configs.Get().Port)
}
