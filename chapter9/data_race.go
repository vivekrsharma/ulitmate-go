package chapter9

import (
	"fmt"
	"sync"
)

var counter int
var lock sync.Mutex
func readAndIncrement() {
	for i:=0; i<2;i++ {
		// The following four lines of code are bad.
		// It is very easy to manifest synchronization issue because
		// fmt. causes a context switch and hence the data is bad.
		//value := counter
		//fmt.Println(value)
		//value++
		//counter = value

		// A good way to go about this is using Atomic instructions.
		// We also need to change counter to be more precision based int to int32 or int64.
		//atomic.AddInt64(&counter, 1)

		lock.Lock()
		value := counter
		// fmt.Println(value) // This isn't good. Try to reduce the work inside critical section.
		value++
		counter = value
		lock.Unlock()
	}
	fmt.Println(counter)
}

func main3() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		incrementCounter()
		wg.Done()
	}()

	go func() {
		incrementCounter()
		wg.Done()
	}()

	wg.Wait()
}
