package controller

import (
	"net/http"
	"sysemp_travel/usecase"

	"sysemp_travel/model"

	"github.com/gin-gonic/gin"
)

type accountToPayController struct {
	AccountToPayUseCase usecase.AccountToPayUseCase
}

func NewAccountToPayController(accountToPayUseCase usecase.AccountToPayUseCase) accountToPayController {
	return accountToPayController{
		AccountToPayUseCase: accountToPayUseCase,
	}
}

func (c *accountToPayController) CreateAccountToPay(ctx *gin.Context) {
	var accountToPay model.AccountToPay

	typ := ctx.Param("type")

	err := ctx.BindJSON(&accountToPay)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.AccountToPayUseCase.CreateAccountToPay(ctx.Request.Context(), typ, accountToPay)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Account to pay created successfully"})
}
