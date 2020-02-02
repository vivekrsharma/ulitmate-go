package chapter5

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) String() string {
	return fmt.Sprintf("The user has name %s with email %s", u.name, u.email)
}

func main() {
	u := user{
		name:  "vrs",
		email: "abc@abc.com",
	}
	fmt.Println(u)
	fmt.Println(&u)
}
