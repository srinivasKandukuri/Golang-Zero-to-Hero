package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter with RWMutex for thread-safe operations
type Counter struct {
	mu    sync.RWMutex
	value int
}

// Increment safely increases the counter
func (c *Counter) Increment() {
	c.mu.Lock()         // Acquire write lock (exclusive access)
	defer c.mu.Unlock() // Release the lock when done
	c.value++
}

// GetValue safely reads the counter
func (c *Counter) GetValue() int {
	c.mu.RLock()         // Acquire read lock (shared access)
	defer c.mu.RUnlock() // Release the lock when done
	return c.value
}

func main() {
	counter := &Counter{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	time.Sleep(2 * time.Second) // Wait for goroutines to finish
	fmt.Println("Final Counter Value:", counter.GetValue())
}

/*
4. Example: Without sync.RWMutex (Data Race Issue)

package main

import (
	"fmt"
	"time"
)

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) GetValue() int {
	return c.value
}

func main() {
	counter := &Counter{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Increment() // Race condition here
			}
		}()
	}

	time.Sleep(2 * time.Second) // Wait for goroutines to finish
	fmt.Println("Final Counter Value:", counter.GetValue())
}

*/
/*
The sync.RWMutex in Go is a read-write mutex (mutual exclusion lock) that provides a way to safely manage concurrent access to shared data.
It is part of Go's sync package and is used to prevent data races.

✅ 1. What Is a Mutex?
A mutex (short for mutual exclusion) is a synchronization primitive used to protect shared resources from
being accessed by multiple goroutines at the same time.

In simple terms:

Lock: Blocks other goroutines from accessing the resource.
Unlock: Allows other goroutines to access the resource.
✅ 2. What Is sync.RWMutex?
sync.RWMutex is a read-write mutex that allows:

Multiple readers to access the resource simultaneously.
Only one writer to access the resource, while blocking all readers.
👉 It is more efficient than a standard sync.Mutex when you have many reads but few writes.


*/
