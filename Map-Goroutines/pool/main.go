package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"time"
)

type fakeConn struct {
	id int
}

func (f *fakeConn) Close() error {
	fmt.Printf("close fakeConn %d\n", f.id)
	return nil
}

var idCounter = 0

func createCon(ctx context.Context) (io.Closer, error) {
	idCounter++
	fmt.Println("create con", idCounter)
	return &fakeConn{id: idCounter}, nil
}

func main() {
	ctx := context.Background()
	myPool, _ := NewPool(createCon, 5)

	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	g, gtx := errgroup.WithContext(ctx)
	for i := 0; i < 10; i++ {
		i := i
		g.Go(func() error {
			conn, err := myPool.Acquire(gtx)
			if err != nil {
				fmt.Println("Failed to acquire pool", err)
				return err
			}
			fmt.Printf("Worker %d got connection\n", i)

			time.Sleep(1 * time.Second) // use it
			err = myPool.Release(conn)
			if err != nil {
				fmt.Println("Failed to release pool", err)
			}
			fmt.Printf("Worker %d returned connection\n", i)
			return err
		})

	}

	err := g.Wait()
	if err != nil {
		fmt.Println("Failed to acquire pool", err)
	}

}
