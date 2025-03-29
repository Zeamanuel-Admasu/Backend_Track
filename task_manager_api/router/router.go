package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/controllers"
	"github.com/zaahidali/task_manager_api/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	router.POST("/promote", middleware.Authenticate(), middleware.Authorize("admin"), controllers.Promote)

	taskRoutes := router.Group("/tasks")
	taskRoutes.Use(middleware.Authenticate()) // All /tasks routes need authentication

	taskRoutes.GET("/", controllers.GetTasks)
	taskRoutes.GET("/:id", controllers.GetTaskById)

	taskRoutes.Use(middleware.Authorize("admin")) // Only admins beyond this point
	taskRoutes.POST("/", controllers.CreateTask)
	taskRoutes.PUT("/:id", controllers.UpdateTask)
	taskRoutes.DELETE("/:id", controllers.DeleteTask)

	return router
}
