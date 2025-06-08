package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userId", 14)

	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	userId := ctx.Value("userId").(int)

	if userId == 14 {
		// do
		fmt.Println("got the userid")
	}
}
