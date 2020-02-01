package integrity

import "fmt"

type notifier interface {
	notify()
}
type user struct {
	name string
	email string
}

func ( u *user) notify() {
	fmt.Printf(" User has name %s and email %s", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user {
		name : "vrs",
		email : "abc@abc.com",
	}
	//sendNotification(u)  .. This is a compile error. We cannot send value in a pointer semantics receiver method.
	// It has following issues -
	// Not all values have addresses.( This is minor issue.) eg. duration(42).notify()
	// The bigger problem is when you use pointer semantics, you can only share the receiver.
	// We cannot create a copy of receiver shared using pointer semantics.
	sendNotification(&u)
}