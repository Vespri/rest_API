package routers

import (
	"github.com/gin-gonic/gin"

	"api_go/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/create", controllers.CreateData)

	router.PUT("/update/:orderID", controllers.UpdateData)

	router.GET("/get_all", controllers.GetAllData)

	router.GET("/get/:orderID", controllers.GetDataByID)

	router.DELETE("/delete/:orderID", controllers.DeleteData)

	return router
}
