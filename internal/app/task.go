package task

type TaskSet struct {
	n   int
	d   float64
	n1  float64
	I   float64
	TTL float64
}

type TaskGet struct {
	number          int
	status          string
	n               int
	d               float64
	n1              float64
	I               float64
	TTL             float64
	iteration       int
	timeOfSetTask   float64
	timeOfStartTask float64
	timeOdEndTask   float64
}
