package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *Handler) getTaskByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id params")
		return
	}

	task, err := hdl.tasksService.Get(ctx, id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, task)
}
