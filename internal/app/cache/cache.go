package cache

import (
	"sync"

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

func (c *Cache) Set(task *task.Task) int {

	// var expiration int64

	// // Если продолжительность жизни равна 0 - используется значение по-умолчанию
	// if duration == 0 {
	// 	duration = c.defaultExpiration
	// }

	// // Устанавливаем время истечения кеша
	// if duration > 0 {
	// 	expiration = time.Now().Add(duration).UnixNano()
	// }

	c.Lock()

	defer c.Unlock()

	c.tasks[c.iterator] = task

	defer c.IncIterator()

	return c.iterator
}

func (c *Cache) GetAll() interface{} {

	c.RLock()

	defer c.RUnlock()

	//_, found := c.tasks[key]

	// ключ не найден
	// if !found {
	// 	return nil, false
	// }

	// Проверка на установку времени истечения, в противном случае он бессрочный
	// if item.Expiration > 0 {

	// 	// Если в момент запроса кеш устарел возвращаем nil
	// 	if time.Now().UnixNano() > item.Expiration {
	// 		return nil, false
	// 	}

	// }

	return c.tasks
}
