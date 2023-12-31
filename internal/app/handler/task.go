package handler

import (
	"net/http"

	"github.com/WORUS/arithmetic-progression/internal/app/task"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, h.services.GetTasks())
}

func (h *Handler) SetTask(c *gin.Context) {
	var tsk task.TaskInput

	if err := c.BindJSON(&tsk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tsk.I < 0 || tsk.TTL < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time cannot be less than 0"})
		return
	}

	h.services.SetTaskInQueue(tsk)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "task successfully queued",
	})
}
