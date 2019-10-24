package router

import "github.com/gin-gonic/gin"
import . "controllers"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/task", AddTaskApi)

	router.GET("/tasks", GetTasksApi)

	//router.GET("/task/:id", GetTaskApi)

	//router.PUT("/task/:id", ModTaskApi)

	//router.DELETE("/task/:id", DelTaskApi)

	return router
}
