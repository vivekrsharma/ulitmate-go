package main

import "fmt"

type reader interface {
	Read(b []byte) (int, error)
}

func main() {
	fmt.Println("This is reader interface")
	r := reader{}
	b := make([]byte,10,10)
	r.Read(b)
}

