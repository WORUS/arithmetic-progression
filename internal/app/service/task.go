package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/WORUS/arithmetic-progression/internal/app/task"
)

func (s *Service) SetTaskInQueue(ctx context.Context, tsk task.TaskInput) error {
	task := task.Task{
		Number:  uint(len(s.queue)),
		Status:  "waiting in queue",
		N:       tsk.N,
		D:       tsk.D,
		N1:      tsk.N1,
		I:       tsk.I,
		TTL:     tsk.TTL,
		SetTime: time.Now().UTC(),
	}
	s.cache.Set(&task)
	//TODO: limiter
	s.queue <- &task
	return nil
	// select {
	// case <-ctx.Done():
	// 	return ctx.Err()
	// case s.queue <- tsk:

	// 	return nil
	// }
}

func (s *Service) QueueListener(ctx context.Context) error {
	for {

		// select {
		// case s.goroutines <- true:
		// 	go s.StartTask(<-s.queue)
		// 	fmt.Printf("Запущено %d задач", len(s.goroutines))
		// }

		select {
		case <-ctx.Done():
			return ctx.Err()
		case tsk := <-s.queue:
			select {
			case <-ctx.Done():
				return ctx.Err()
			case s.goroutines <- true:
				go s.StartTask(tsk)
			}

		}
	}

}

func (s *Service) StartTask(tsk *task.Task) {
	//s.goroutines <- true

	value := fmt.Sprintf("%fs", tsk.I)
	interval, err := time.ParseDuration(value)
	if err != nil {
		log.Fatal(err.Error())
	}
	ticker := time.NewTicker(interval)

	func() {
		a := make([]float64, tsk.N)
		a[0] = tsk.N1

		iter := 1

		tsk.Status = "in process"
		tsk.StartTime = time.Now()
		tsk.Iteration = uint(iter)

		for t := range ticker.C {
			if iter < len(a) {
				a[iter] = a[iter-1] + tsk.D
				iter++

				tsk.Iteration = uint(iter)

				fmt.Println("tick at ", t.UTC())
				continue
			}

			tsk.EndTime = time.Now()
			tsk.Status = "completed"

			return
		}
	}()

	ticker.Stop()
	<-s.goroutines
}

func (s *Service) GetTasks() interface{} {
	return s.cache.GetAll()
}
