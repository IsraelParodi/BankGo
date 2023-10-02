package router

import (
	"github.com/gin-gonic/gin"
	account "github.com/israelparodi/bankgo/api/accounts"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("book", Controllers.ListBook)
	// 	v1.POST("book", Controllers.AddNewBook)
	// 	v1.GET("book/:id", Controllers.GetOneBook)
	// 	v1.PUT("book/:id", Controllers.PutOneBook)
	// 	v1.DELETE("book/:id", Controllers.DeleteBook)
	// }

	router.POST("/accounts", account.CreateAccount)

	return router
}
