package v2

import "sync"

type Counter struct {
	value int

	// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	sync.Mutex
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
