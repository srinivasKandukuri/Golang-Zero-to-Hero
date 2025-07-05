package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	orderChan := make(chan *Order, 20)
	processedChan := make(chan *Order, 20)
	go func() {
		defer wg.Done()
		for _, order := range generateOrders(20) {
			orderChan <- order
		}
		close(orderChan)
		fmt.Println("Done with generating orders")
	}()

	go processOrder(orderChan, processedChan, &wg)

	go func() {
		defer wg.Done()
		for {
			select {
			case processOrder, ok := <-processedChan:
				if !ok {
					fmt.Printf("Processing orders closed")
					return
				}
				fmt.Printf("Processing order %s :  %d\n", processOrder.Status, processOrder.ID)
			case <-time.After(100 * time.Millisecond):
				fmt.Printf("Processing orders timeout")
				return
			}
		}
	}()

	wg.Wait()

	fmt.Println("Orders completed.")
}

func processOrder(inChan <-chan *Order, outChan chan<- *Order, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(outChan)
	}()
	for order := range inChan {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		order.Status = "processed"
		outChan <- order
	}
}
func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i + 1, Status: "Pending"}
	}
	return orders
}

func updateOrderStatus(order *Order) {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)

	status := []string{"Processing", "Shipped", "Delivered"}[rand.Intn(3)]

	order.Status = status
	fmt.Printf("Updating order %d status to %s\n", order.ID, order.Status)

}
