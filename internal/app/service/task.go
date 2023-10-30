package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/WORUS/arithmetic-progression/internal/app/task"
)

func (s *Service) SetTaskInQueue(ctx context.Context, tsk task.TaskInput) error {

	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.queue <- tsk:
		//fmt.Print(<-s.queue)
		return nil
	}
}

func (s *Service) QueueListener(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case tsk := <-s.queue:
			go s.StartTask(tsk)
		}
	}

}

func (s *Service) StartTask(tsk task.TaskInput) task.Task {
	var task task.Task

	a := make([]float64, tsk.N)
	a[0] = tsk.N1

	// for i := 1; i < len(a); i++ {
	// 	a[i] = a[i-1] + tsk.D
	// }

	value := fmt.Sprintf("%fs", tsk.I)
	interval, err := time.ParseDuration(value)
	if err != nil {
		log.Fatal(err.Error())
	}
	ticker := time.NewTicker(interval)

	func() {
		iter := 1
		fmt.Println("Progression started")
		for t := range ticker.C {
			if iter < tsk.N {
				a[iter] = a[iter-1] + tsk.D
				iter++
				fmt.Println("tick at ", t.UTC())
				continue
			}
			return
		}
	}()
	fmt.Print(a)
	ticker.Stop()

	return task
}
