package sync

import "sync"

type Counter struct {
	// Mutex allows us to add locks to our data
	mu    sync.Mutex
	value int
}

// to force to be used instead of create a copy of counter because of Mutex
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
