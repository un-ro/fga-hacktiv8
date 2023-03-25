package routers

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", CreateBook)

	return router
}
