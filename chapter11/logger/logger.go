package logger

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

func New(w io.Writer, capacity int) *Logger {
	l := Logger{
		ch: make(chan string, capacity),
		wg: sync.WaitGroup{},
	}
	l.wg.Add(1)

	go func() {
		for l := range l.ch {
			fmt.Fprint(w, l)
		}
		l.wg.Done()
	}()

	return &l
}

func (l *Logger) Close() {
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Write(line string) {
	select {
	case l.ch <- line:
	default:
		fmt.Println("line dropped.")
	}
}
