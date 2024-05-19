package handler

import (
	"Skillfactory_task_30.8.1/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	tasksService service.TasksService
}

func NewHandler(tasksService service.TasksService) *Handler {
	return &Handler{tasksService: tasksService}
}

func (hdl *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		tasks := api.Group("/tasks")
		{
			tasks.GET("/:id", hdl.getTaskByID)
		}
	}

	return router
}
