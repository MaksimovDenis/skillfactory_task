package handler

import (
	"net/http"

	"Skillfactory_task_30.8.1/internal/models"
	"github.com/gin-gonic/gin"
)

func (hdl *Handler) createTask(ctx *gin.Context) {
	var task models.CreateTasks

	if err := ctx.BindJSON(&task); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid request body")
		return
	}

	taskID, err := hdl.tasksService.CreateTask(ctx.Request.Context(), &task)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"task_id": taskID})
}

func (hdl *Handler) getTasks(ctx *gin.Context) {

	task, err := hdl.tasksService.GetTasks(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (hdl *Handler) getTaskByAuthorName(ctx *gin.Context) {

	authorName := ctx.Param("authorName")
	if authorName == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "Author name cannot be empty")
		return
	}

	task, err := hdl.tasksService.GetTaskByAuthorName(ctx, authorName)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (hdl *Handler) getTaskByLabel(ctx *gin.Context) {

	label := ctx.Param("label")
	if label == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "Author name cannot be empty")
		return
	}

	task, err := hdl.tasksService.GetTaskByLabel(ctx, label)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, task)
}
