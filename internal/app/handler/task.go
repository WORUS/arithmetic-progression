package handler

import (
	"context"
	"net/http"

	"github.com/WORUS/arithmetic-progression/internal/app/task"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetTask(c *gin.Context) {
	var tsk task.TaskInput
	ctx := context.Background()
	if err := c.BindJSON(&tsk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.services.SetTaskInQueue(ctx, tsk)
	c.JSON(http.StatusOK, tsk)
}