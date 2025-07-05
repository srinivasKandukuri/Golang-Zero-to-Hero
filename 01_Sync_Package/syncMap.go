package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// to store day using sync.Map
	m.Store("name", "sk")
	m.Store("age", 30)

	if value, ok := m.Load("name"); ok {
		fmt.Println("Name", value)
	}

	// LoadOrStore- Useful for caching
	actual, loaded := m.LoadOrStore("city", "Hyderabad")
	fmt.Println("City", actual, "already present ", loaded)

	m.Delete("age")

	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

/*
🧵 sync.Map in Go
sync.Map is a concurrent-safe map provided by the Go sync package. It is designed for use cases where multiple goroutines need to read and write to the same map concurrently without manual locking.

✅ When to Use sync.Map
High-Concurrency Access: When multiple goroutines need to access and modify a shared map.
Read-Mostly Workloads: Ideal when most operations are reads rather than writes.
Caching: Useful for caching frequently accessed but rarely updated data.
Dynamic State Management: For storing dynamic configurations shared across goroutines.
📚 Basic Operations
Here are the main methods of sync.Map:

Store(key, value) – Add or update a key-value pair.
Load(key) – Retrieve a value by key (returns value, ok).
LoadOrStore(key, value) – Load if exists, otherwise store (returns actual, loaded).
Delete(key) – Remove a key-value pair.
Range(func(key, value) bool) – Iterate over all entries.


🚀 When NOT to Use sync.Map
Write-Heavy Workloads – For frequent updates, sync.Mutex with a regular map is more efficient.
Fixed Key Set – If keys are known and static, a locked map is simpler and faster.
*/
