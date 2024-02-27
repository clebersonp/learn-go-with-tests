package v2

import "sync"

type Counter struct {
	value int

	// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	// Use channels when passing ownership of data
	// Use mutexes for managing state (this case)
	// Don't use embedding because it's convenient as anonymous field without its name because the field will be
	// exported and accessible for any one. The file will become public for every one
	mu sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
