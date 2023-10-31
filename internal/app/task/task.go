package task

import (
	"time"
)

type TaskInput struct {
	N   uint    `json:"n" binding:"required"`
	D   float64 `json:"d" binding:"required"`
	N1  float64 `json:"n1" binding:"required"`
	I   float64 `json:"I" binding:"required"`
	TTL float64 `json:"TTL" binding:"required"`
}

type Task struct {
	Number    uint      `json:"number"` //номер в очереди
	Status    string    `json:"status"`
	N         uint      `json:"n"`
	D         float64   `json:"d"`
	N1        float64   `json:"n1"`
	I         float64   `json:"I"`
	TTL       float64   `json:"TTL"`
	Iteration uint      `json:"iteration"` //текущая итерация
	SetTime   time.Time `json:"setTime"`   //время занесения задачи в очередь
	StartTime time.Time `json:"startTime"` //время старта задачи
	EndTime   time.Time `json:"endTime"`   //время окончания задачи(если выполнена)
	Key       int
}
