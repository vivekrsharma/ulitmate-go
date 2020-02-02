package main

import "fmt"

type user struct {
	name  string
	email string
}

func main() {
	u1 := createV1User()
	fmt.Print("u1")
	u2 := createV2User()
	fmt.Println("&u1", &u1, "\nu1 ", u1, "\n&u2", &u2, "\nu2", u2)
}

func createV1User() user {
	u := user{
		name:  "abc",
		email: "abc@abc.com",
	}
	return u
}

func createV2User() *user {
	u := user{
		name:  "abc",
		email: "abc@abc.com",
	}
	return &u
}
