package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
