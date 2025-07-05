package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go initDb(i, &wg)
	}
	wg.Wait()

}

func initDb(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(initialize)
	fmt.Println("worker", id, "stated")
}
func initialize() {
	fmt.Println("init db ")
}

/*
🔹 Explanation
once.Do(initialize): Ensures initialize() runs only once, regardless of how many goroutines invoke it.
Only the first goroutine that calls once.Do() will execute initialize().
🔹 Best Practices
✅ Use for initialization logic (e.g., database connection, config loading).
✅ Avoid placing heavy operations inside once.Do(), as it blocks all goroutines until completion.
*/
