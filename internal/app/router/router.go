package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/widcha/openidea-marketplace/internal/app"
	"github.com/widcha/openidea-marketplace/internal/app/modules/health"
	userhandler "github.com/widcha/openidea-marketplace/internal/app/modules/user/handler"
	"github.com/widcha/openidea-marketplace/internal/middleware"
	"github.com/widcha/openidea-marketplace/internal/pkg"
)

type Router struct {
	router     gin.IRouter
	datasource *pkg.Datasource
	container  *app.Container
}

func NewRouter(router gin.IRouter, datasource *pkg.Datasource, container *app.Container) *Router {
	return &Router{
		router:     router,
		datasource: datasource,
		container:  container,
	}
}

func (h *Router) RegisterRouter() {
	h.router.Use(middleware.PrometheusMiddleware())
	h.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Welcome to openidea-marketplace",
		})
	})
	h.router.GET("/health", health.GetHealthCheckHandler(h.container.HealthCheckUsecase))
	h.router.GET("/prometheus", gin.WrapH(promhttp.Handler()))

	v1 := h.router.Group("/v1")

	// User
	user := v1.Group("/user")
	user.POST("/register", userhandler.SignupHandler(h.container.UserSignupUsecase))
	user.POST("/login", userhandler.SigninHandler(h.container.UserSigninUsecase))

}
