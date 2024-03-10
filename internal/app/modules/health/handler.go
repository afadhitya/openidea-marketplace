package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealthCheckHandler(usecase Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := usecase.HealthCheck(c.Copy().Request.Context())
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "failed",
				"status":  resp.Status,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"status":  resp.Status,
		})
	}
}
