package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/WORUS/arithmetic-progression/internal/app/task"
)

func (s *Service) SetTaskInQueue(tsk task.TaskInput) error {
	task := task.Task{
		Number:  uint(len(s.queue)) + 1,
		Status:  "waiting in queue",
		N:       tsk.N,
		D:       tsk.D,
		N1:      tsk.N1,
		I:       tsk.I,
		TTL:     tsk.TTL,
		SetTime: time.Now().Format(time.DateTime),
	}

	s.cache.Set(&task)

	s.queue = append(s.queue, &task)

	s.qready <- true

	return nil
}

func (s *Service) queueIndexResfresher() {
	for i := range s.queue {
		s.queue[i].Number = uint(i + 1)
	}
}

func (s *Service) QueueListener(ctx context.Context) {
	for {

		select {
		case <-ctx.Done():
			return
		case <-s.qready:
			select {

			case <-ctx.Done():
				return

			case s.goroutines <- true:
				go s.startTask(s.queue[0])

				s.queue[0].Number = 0
				s.queue = s.queue[1:]

				go s.queueIndexResfresher()

			}
		}

	}
}

func (s *Service) startTask(tsk *task.Task) {

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
		tsk.StartTime = time.Now().Format(time.DateTime)
		tsk.Iteration = uint(iter)

		for range ticker.C {
			if iter < len(a) {
				a[iter] = a[iter-1] + tsk.D
				iter++

				tsk.Iteration = uint(iter)

				continue
			}

			tsk.EndTime = time.Now().Format(time.DateTime)
			tsk.Status = "completed"

			return
		}
	}()

	ticker.Stop()

	go s.cache.TaskCleaner(tsk)

	<-s.goroutines
}

func (s *Service) GetTasks() interface{} {
	return s.cache.GetAll()
}
