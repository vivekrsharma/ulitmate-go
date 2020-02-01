package composition

import "fmt"

type speak interface {
	Speak()
}

type myDog struct {
	name string
	isMammal bool
	packFactor int
}

func (d myDog) Speak() {
	fmt.Println("Name %s is %q and packaFactor %d",d.name,d.isMammal,d.packFactor)
}

type myCat struct {
	name string
	isMammal bool
	climbFactor int
}

func (c myCat) Speak() {
	fmt.Println("Name %s is %q and packaFactor %d",c.name,c.isMammal,c.climbFactor)
}

func main() {
	slice := []speak {
		myCat{
			name:        "kitty",
			isMammal:    true,
			climbFactor: 10,
		},
		myDog{
			name:       "doggy",
			isMammal:   true,
			packFactor: 100,
		},
	}

	fmt.Println(slice)
}