package chapter9

import (
	"fmt"
	"sync"
)


func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	lookup := make(map[string]int)
	go func() {
		for i:=0;i<1000;i++ {
			lookup["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i:=0;i<1000;i++ {
			lookup["B"]++
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(lookup)
}


