package main

import "fmt"

func main() {
	var data []string
	data = []string{}
	// Consider doing the following since we  know the size.
	//data := make([]string, 0,100000)
	for record := 0; record<=100000;record++ {
		lastCap := cap(data)
		str := fmt.Sprint("%s",record)
		data = append(data, str)

		if lastCap != cap(data) {
			percent := (float64(cap(data) - lastCap)) / float64(lastCap) * 100.0
			fmt.Printf(" lastCapacity [%d] currentCapacity [%d] percentChange [%f] \n", lastCap, cap(data), percent )
		}
	}
}

