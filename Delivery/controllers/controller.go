package controllers

import (
	"net/http"
	// "task-manager/Domain"
	"task-manager/Usecases"

	"github.com/gin-gonic/gin"
)

// ================== TASK CONTROLLER ==================
type TaskController struct {
	usecase Usecases.TaskUsecase
}

func NewTaskController(u Usecases.TaskUsecase) *TaskController {
	return &TaskController{usecase: u}
}

// GET /tasks
func (ctrl *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := ctrl.usecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// GET /tasks/:id
func (ctrl *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := ctrl.usecase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// POST /tasks
func (ctrl *TaskController) CreateTask(c *gin.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := ctrl.usecase.CreateTask(req.Title, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

// PUT /tasks/:id
func (ctrl *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := ctrl.usecase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	updated, err := ctrl.usecase.UpdateTask(task, req.Title, req.Description, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DELETE /tasks/:id
func (ctrl *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.usecase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

// ================== USER CONTROLLER ==================
type UserController struct {
	usecase Usecases.UserUsecase
}

func NewUserController(u Usecases.UserUsecase) *UserController {
	return &UserController{usecase: u}
}

// POST /register
func (ctrl *UserController) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.usecase.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return safe JSON
	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID(),
		"username": user.Username(),
		"role":     user.Role(),
	})
}

// POST /login
func (ctrl *UserController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.usecase.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// PUT /users/:username/promote
func (ctrl *UserController) Promote(c *gin.Context) {
	username := c.Param("username")
	if err := ctrl.usecase.Promote(username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}