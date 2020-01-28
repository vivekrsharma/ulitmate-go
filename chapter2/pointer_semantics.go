package main

import "fmt"

func main() {
	counter := 10
	increment_referene(&counter)
	fmt.Println(counter)
}

func increment_referene(inc *int) {
	*inc++
	fmt.Println(*inc)
}