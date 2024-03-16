package bankaccounthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bankaccountusecase "github.com/widcha/openidea-marketplace/internal/app/modules/bankaccount/usecase/savebankaccount"
)

func SaveBankAccountHandler(inport bankaccountusecase.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req bankaccountusecase.InportRequest

		c.BindJSON(&req)

		resp, err := inport.Execute(c.Copy().Request.Context(), req)
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
			"data":    resp,
		})
	}
}
