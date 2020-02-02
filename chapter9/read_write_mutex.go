package chapter9

import (
"fmt"
"sync"
	"sync/atomic"
)

var shared int
var rwLock sync.RWMutex
var readCount int64
func incrementCounter() {
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

		rwLock.Lock()
		value := shared
		// fmt.Println(value) // This isn't good. Try to reduce the work inside critical section.
		value++
		shared = value
		fmt.Printf("Writer is inside %d. ReadCount %d \n",shared,atomic.LoadInt64(&readCount))
		rwLock.Unlock()
	}
	fmt.Println(shared)
}


func readCounter() {
	for i:=0; i<2;i++ {

		rwLock.RLock()
		atomic.AddInt64(&readCount,1)
		fmt.Printf("The read value is %d and readCount is %d \n",shared, atomic.LoadInt64(&readCount))
		atomic.AddInt64(&readCount,-1)

		rwLock.RUnlock()
	}
}
func main2() {
	var wg sync.WaitGroup
	wg.Add(9)

	go func(){
		incrementCounter()
		wg.Done()
	}()

	for i:=0;i<8;i++ {
		go func() {
			readCounter()
			wg.Done()
		}()
	}

	wg.Wait()
}

