package router

import (
	"github.com/gin-gonic/gin"
	account "github.com/israelparodi/bankgo/api/account"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/account", account.CreateAccount)
		v1.GET("/account/:id", account.GetAccount)
	}

	return router
}
