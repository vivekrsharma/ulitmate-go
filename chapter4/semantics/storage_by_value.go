package semantics

import "fmt"

type user struct {
	name string
}

func (u user) Print() {
	fmt.Println("My name is %s",u.name)
}

type printer interface {
	Print()
}

func main() {
	u := user{
		name:"vrs",
	}

	slice := []printer {
		u,
		&u,
	 }
	 fmt.Println("Before..")
	for i:=0;i<len(slice);i++ {
		slice[i].Print()
	}
	 u.name="CHANGED"
	fmt.Println("After..")
	 // Since we used value semantics for storing the user in slice at index 0,
	 // the changes won't be reflected in the index 0.
	for i:=0;i<len(slice);i++ {
		slice[i].Print()
	}

}
