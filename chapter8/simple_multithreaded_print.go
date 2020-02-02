package chapter8

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func lowercase() {
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c", ch)
	}
	fmt.Println("Done")
}

func uppercase() {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%c", ch)
	}
	fmt.Println("Done")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 5; i++ {
			lowercase()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			uppercase()
		}
		wg.Done()
	}()

	wg.Wait()
}
