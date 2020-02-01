package chapter8

import (
	"fmt"
	"runtime"
	"sync"
)

func primeNumber(n int, name string) {
	for i:=1; i<=n;i++ {
		isPrime := true
		for j:=2;j<i;j++ {
			if i %j == 0 {
				isPrime = false
				break;
			}
		}
		if isPrime {
			fmt.Printf("%s: %d \n", name, i)
		}
	}
}

func init() {
	runtime.GOMAXPROCS(2)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		primeNumber(5000, "A")
		wg.Done()
	}()

	go func (){
		primeNumber(5000, "B")
		wg.Done()
	}()

	wg.Wait()
}
