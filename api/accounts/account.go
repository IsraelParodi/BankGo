package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/israelparodi/bankgo/api/helpers"
	"github.com/israelparodi/bankgo/config"
	db "github.com/israelparodi/bankgo/db/sqlc/queries"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	arg := db.CreateAccountParams{Owner: req.Owner, Currency: req.Currency, Balance: 0}
	account, err := config.Queries.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, account)
}
