package handler

import (
	"net/http"

	"github.com/WORUS/arithmetic-progression/internal/app/task"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetTask(c *gin.Context) {
	var tsk task.TaskInput

	if err := c.BindJSON(&tsk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
