package router

import (
	"task_manager/controllers"
	"task_manager/middleware"
	"github.com/gin-gonic/gin"
	
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/tasks", controllers.GetTasks)
		auth.GET("/tasks/:id", controllers.GetTask)

		admin := auth.Group("/")
		admin.Use(middleware.AdminOnly())
		{
			admin.POST("/tasks", controllers.CreateTask)
			admin.PUT("/tasks/:id", controllers.UpdateTask)
			admin.DELETE("/tasks/:id", controllers.DeleteTask)
			admin.POST("/promote/:username", controllers.PromoteUser)
		}
	}
	return r
}
