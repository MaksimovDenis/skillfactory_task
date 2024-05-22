package handler

import (
	"github.com/gin-gonic/gin"
	"skillfactory_task_30.8.1/internal/service"
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
