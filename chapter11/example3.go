package main

import (
	"fmt"
	logger2 "github.com/vivekrsharma/ulitmate-go/chapter11/logger"
	"os"
	"os/signal"
	"time"
)

// device helps simulate io.Writer to write logs.
type device struct {
	problem bool
}

func (d *device) Write(p []byte) ( n int, err error) {
	if d.problem {
		time.Sleep(1 * time.Second)
	}
	fmt.Println(string(p))
	return len(p), nil
}

func main() {

	const grps=10

	var d device
	logger := logger2.New(&d,grps)
	defer logger.Close()

	for i:=0;i<grps;i++ {
		go func(id int) {
			// Endlessly keep working and keep logging.
			for {
				// simulate some work.
				time.Sleep(10 * time.Millisecond)
				logger.Write(fmt.Sprintf("%d: logging data", id))
			}
		}(i)
 	}

	// To be able to capture the simulated disk blocking, we want
	// to be able to capture the interrupts.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan

		// keep toggling the device status.
		// this contains a data race but it is ok to simulate.
		d.problem = !d.problem
	}
}
