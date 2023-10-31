package cache

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/WORUS/arithmetic-progression/internal/app/task"
)

type Cache struct {
	sync.RWMutex
	tasks    map[int]*task.Task
	iterator int
}

func NewCache() *Cache {
	tasks := make(map[int]*task.Task)
	cache := Cache{
		tasks:    tasks,
		iterator: 0,
	}
	return &cache
}

func (c *Cache) IncIterator() {
	c.iterator++
}

func (c *Cache) Set(task *task.Task) {

	c.Lock()

	defer c.Unlock()

	c.tasks[c.iterator] = task

	task.Key = c.iterator

	defer c.IncIterator()

}

func (c *Cache) GetAll() interface{} {

	c.RLock()

	defer c.RUnlock()

	return c.tasks
}

func (c *Cache) TaskCleaner(tsk *task.Task) {

	value := fmt.Sprintf("%fs", tsk.TTL)
	interval, err := time.ParseDuration(value)
	if err != nil {
		log.Fatal(err.Error())
	}

	tickerTTL := time.NewTicker(interval)

	<-tickerTTL.C

	delete(c.tasks, tsk.Key)
}
