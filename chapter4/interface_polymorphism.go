package main

import "fmt"

type reader interface {
	read(b []byte) (int, error)
}

type pipe struct {}
func (p pipe) read(b []byte) (int, error) {
	return 10, nil
}

type file struct {
	name string
}

func (f file) read(b []byte) (int, error) {
	return len(b), nil
}

func retrieve ( r reader) {
	b := make([]byte,10,10)
	i, err := r.read(b)
	if err!= nil {
		panic(err)
	}
	fmt.Printf("Read %d amount of bytes", i)
}

func main() {
	fmt.Println("This is reader interface")
	p := pipe{}
	f := file{name:"abc.txt"}

	retrieve(p)
	retrieve(f)
}

