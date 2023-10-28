package task

import "time"

type TaskInput struct {
	N   int     `json:"n" binding:"required"`
	D   float64 `json:"d" binding:"required"`
	N1  float64 `json:"n1" binding:"required"`
	I   float64 `json:"I" binding:"required"`
	TTL float64 `json:"TTL" binding:"required"`
}

type Task struct {
	number    int `json:"number"`
	status    string
	N         int       `json:"n" binding:"required"`
	D         float64   `json:"d" binding:"required"`
	N1        float64   `json:"n1" binding:"required"`
	I         float64   `json:"I" binding:"required"`
	TTL       float64   `json:"TTL" binding:"required"`
	iteration int       `json:"iteration"`
	setTime   time.Time `json:"setTime"`
	startTime time.Time `json:"startTime"`
	endTime   time.Time `json:"endTime"`
}
