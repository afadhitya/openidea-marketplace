package userhandler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	userusecase "github.com/widcha/openidea-marketplace/internal/app/modules/user/usercase/signup"
)

func SignupHandler(inport userusecase.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req userusecase.InportRequest

		c.BindJSON(&req)

		resp, err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			if strings.Contains(err.Error(), "users_un") {
				c.JSON(http.StatusBadRequest, gin.H{
					"Success": false,
					"Message": "unsuccessfully signup user",
					"data":    "username already registered",
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "successfully signup user",
			"data":    resp,
		})
	}
}
