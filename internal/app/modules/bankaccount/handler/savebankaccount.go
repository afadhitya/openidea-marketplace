package bankaccounthandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	bankaccountusecase "github.com/widcha/openidea-marketplace/internal/app/modules/bankaccount/usecase/savebankaccount"
	"github.com/widcha/openidea-marketplace/internal/pkg/token"
)

func SaveBankAccountHandler(inport bankaccountusecase.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req bankaccountusecase.InportRequest

		tkn, ok := c.Get("user")
		if !ok {
			log.Println("unable to retrieve user token")
			return
		}

		user, ok := tkn.(*token.UserClaims)
		if !ok {
			log.Println("unable read user claims")
			return
		}

		c.BindJSON(&req)

		req.UserID = user.UserID

		err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"Message": "unsuccessfully save bank account",
				"data":    "unsuccessfully save bank account",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "successfully save bank account",
		})
	}
}
