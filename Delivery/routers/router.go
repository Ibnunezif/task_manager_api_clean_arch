package routers

import (
	"task-manager/Delivery/controllers"
	"task-manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes all routes and middleware
func SetupRouter(taskCtrl *controllers.TaskController, userCtrl *controllers.UserController) *gin.Engine {
	r := gin.Default()

	// Public User Routes
	r.POST("/register", userCtrl.Register)
	r.POST("/login", userCtrl.Login)

	// JWT-protected routes
	auth := r.Group("/")
	auth.Use(Infrastructure.AuthMiddleware())

	// Task routes (authenticated users)
	auth.GET("/tasks", taskCtrl.GetAllTasks)
	auth.GET("/tasks/:id", taskCtrl.GetTaskByID)

	// Task routes (admin only)
	admin := auth.Group("/")
	admin.Use(Infrastructure.AdminMiddleware())
	admin.POST("/tasks", taskCtrl.CreateTask)
	admin.PUT("/tasks/:id", taskCtrl.UpdateTask)
	admin.DELETE("/tasks/:id", taskCtrl.DeleteTask)

	// Admin-only user promotion
	admin.PUT("/promote/:username", userCtrl.Promote)

	return r
}