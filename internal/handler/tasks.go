package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *Handler) getTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id params")
		return
	}

	task, err := hdl.taskService.Get()
}
