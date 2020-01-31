package chapter4

import "fmt"

type user struct {
	name string
	email string
}

func (u *user) changeEmail(newEmail string) {
	u.email = newEmail
}

func (u *user) displayUserWithPointerSemantics() {
	fmt.Printf(" %s has email %s \n",u.name,u.email)
}

// Dont ever mix the semantics on the method. This is just for an example.
func (u user) displayUserWithValueSemantics() {
	fmt.Printf(" %s has email %s \n",u.name,u.email)
}

func main() {
	d := user {
		name :"vrs",
		email :"abc",
	}

	f1 := d.displayUserWithPointerSemantics
	f2 := d.displayUserWithValueSemantics

	d.changeEmail("xyz")
	d.displayUserWithPointerSemantics()  // Reflects the new email.
	d.displayUserWithValueSemantics() // Reflect the new email.
	f1() // Since this is pointer semantics, it will reflect the new email.
	f2() // Since this is value semantics, a copy was made and it only refers to the old email.

	users := []user {
		{"abc","abc@abc.com"},
		{"xyz","xyz@xyz.com"},
	}

	// Value semantics of for range. The original users aren't gonna be changed.
	for _,v := range users {
		v.changeEmail("CHANGED")
	}

	for i := range users {
		users[i].displayUserWithValueSemantics()
	}

	for i := range users {
		users[i].changeEmail("CHANGED")
	}

	for i := range users {
		users[i].displayUserWithValueSemantics()
	}
}

//GOPATH=/Users/vsharma/git/go #gosetup
///Users/vsharma/.goenv/versions/1.12.7/bin/go build -o /private/var/folders/2d/8pdylyf94mb8zmx_vsj9pcq00000gp/T/___go_build_method_semantics_go /Users/vsharma/git/go/src/github.com/vivekrsharma/ulitmate-go/chapter4/method_semantics.go #gosetup
///private/var/folders/2d/8pdylyf94mb8zmx_vsj9pcq00000gp/T/___go_build_method_semantics_go #gosetup
// vrs has email xyz
// vrs has email xyz
// vrs has email xyz
// vrs has email abc
// abc has email abc@abc.com
// xyz has email xyz@xyz.com
// abc has email CHANGED
// xyz has email CHANGED