package handler

import (
	"skillfactory_task/internal/service"

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
			tasks.POST("/", hdl.createTask)

			tasks.GET("/", hdl.getTasks)
			tasks.GET("/:id", hdl.getTaskById)
			tasks.PATCH("/", hdl.updateTaskById)
			tasks.DELETE("/:id", hdl.deleteTaskById)

			author := api.Group("/author")
			{
				author.GET("/:authorName", hdl.getTaskByAuthorName)
			}

			label := api.Group("/label")
			{
				label.GET("/:label", hdl.getTaskByLabel)
			}
		}
	}

	return router
}
