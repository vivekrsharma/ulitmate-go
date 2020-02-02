package main

import (
	"fmt"
	"sync"
)

type  Speaker interface {

	Speak() bool
}

type Ben struct {
	name string
	like int
}

func (b *Ben) Speak() bool {
	if b.name != "Ben" {
		fmt.Printf("Ben says, My name is %s.",b.name)
		return false
	}
	return true
}

type Jerry struct {
	name string
	like int
}

func (j *Jerry) Speak() bool {
	if j.name != "Jerry" {
		fmt.Printf("Jerry says, my name is %s", j.name)
		return false
	}
	return true
}

func main() {
	var candidate Speaker

	var wg sync.WaitGroup
	wg.Add(2)
	jerry := Jerry{
		name: "Jerry",
		like: 0,
	}

	ben := Ben{
		name: "Ben",
		like: 0,
	}

	go func() {
		for {
			candidate = &ben
			if !candidate.Speak() {
				panic("Race detected")
			}
		}
		wg.Done()
	}()

	go func() {
		for {
			candidate = &jerry
			candidate.Speak()
			if !candidate.Speak() {
				panic("Race detected")
			}
		}
		wg.Done()
	}()

	wg.Wait()
}