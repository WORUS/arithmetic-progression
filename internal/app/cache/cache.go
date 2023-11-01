package cache

import (
	"fmt"
	"log"
	"sort"
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

	length := len(c.tasks)

	keys := make([]int, 0, length)

	for k := range c.tasks {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	order := make([]task.Task, 0, length)

	for k := range keys {
		order = append(order, *c.tasks[keys[k]])
	}

	return order
}

func (c *Cache) TaskCleaner(tsk *task.Task) {

	value := fmt.Sprintf("%fs", tsk.TTL)
	interval, err := time.ParseDuration(value)
	if err != nil {
		log.Fatal(err.Error())
	}

	tickerTTL := time.NewTicker(interval)

	<-tickerTTL.C

	c.Lock()

	defer c.Unlock()

	delete(c.tasks, tsk.Key)
}
