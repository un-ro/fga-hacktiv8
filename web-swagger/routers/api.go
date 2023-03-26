package routers

import (
	"web-swagger/controllers"
	_ "web-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// Read
	router.GET("/cars/:id", controllers.GetOneCars)

	//Create
	router.POST("/cars", controllers.CreateCars)

	// Read All
	router.GET("/cars", controllers.GetAllCars)

	// Update
	router.PATCH("/cars/:id", controllers.UpdateCars)

	// Delete
	router.DELETE("/cars/:id", controllers.DeleteCars)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
