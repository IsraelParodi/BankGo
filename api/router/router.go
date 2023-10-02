package router

import (
	"github.com/gin-gonic/gin"
	account "github.com/israelparodi/bankgo/api/account"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		accountGroup := v1.Group("/accounts")
		{
			accountGroup.POST("/", account.CreateAccount)
			accountGroup.GET("/:id", account.GetAccount)
			accountGroup.GET("/", account.ListAccounts)
		}
	}

	return router
}
