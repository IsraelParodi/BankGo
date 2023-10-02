package account

import (
	"database/sql"
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

	args := db.CreateAccountParams{Owner: req.Owner, Currency: req.Currency, Balance: 0}
	account, err := config.Queries.CreateAccount(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	account, err := config.Queries.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, helpers.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=5,max=10"`
}

func ListAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	args := db.ListAccountsParams{
		Limit:  int32(req.PageSize),
		Offset: int32(req.PageID-1) * int32(req.PageSize),
	}
	account, err := config.Queries.ListAccounts(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
