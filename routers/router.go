package routers

import (
	"github.com/gin-gonic/gin"

	"api_go/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/create", controllers.CreateData)

	router.PUT("/update/:orderID", controllers.UpdateData)

	// router.GET("/cars_all", controllers.GetCar)

	// router.GET("/cars/:carID", controllers.GetCar)

	// router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
