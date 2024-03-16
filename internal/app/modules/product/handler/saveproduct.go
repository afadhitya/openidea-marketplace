package producthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productusecase "github.com/widcha/openidea-marketplace/internal/app/modules/product/usecase/saveproduct"
)

func SaveProductHandler(inport productusecase.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req productusecase.InportRequest

		c.BindJSON(&req)

		err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"Message": "unsuccessfully product account",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "successfully save product",
		})
	}
}
