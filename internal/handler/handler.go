package handler

import (
	"Skillfactory_task_30.8.1/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.TasksService
}

func NewHandler(service *service.TasksService) *Handler {
	return &Handler{service: service}
}

func (hdl *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		tasks := api.Group("/tasks")
		{
			tasks.GET("/:id")
		}
	}
}
