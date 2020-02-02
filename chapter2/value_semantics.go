package main

import "fmt"

func main() {
	counter := 10
	increment(counter)
	fmt.Println(counter)
}

func increment(inc int) {
	inc++
	fmt.Println(inc)
}
