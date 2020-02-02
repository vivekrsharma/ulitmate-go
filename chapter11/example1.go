package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func contextTimeout() {

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "123"
	}()

	select {
	case <-ch:
		fmt.Println("Task completed successfully.")
	case <-ctx.Done():
		fmt.Println("The action is timedout.")
	}
}

func main() {
	contextTimeout()
}
