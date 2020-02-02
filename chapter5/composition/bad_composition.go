package composition

import "fmt"

//
type Animal struct {
	Name      string
	IsMammmal bool
}

func (a *Animal) Speak() {
	fmt.Printf("My name is %s, it is %t I am a mamalal \n", a.Name, a.IsMammmal)
}

type Dog struct {
	Animal
	Packfactor int
}

func (d *Dog) speak() {
	fmt.Printf(" Woof! My name is %s , it is %t I am a mammal with a pack factor of %d", d.Name, d.IsMammmal, d.Packfactor)
}

type Cat struct {
	Animal
	ClimbFactor int
}

func (c *Cat) speak() {
	fmt.Printf(" Woof! My name is %s , it is %t I am a mammal with a pack factor of %d", c.Name, c.IsMammmal, c.ClimbFactor)
}

// Grouping dogs and cats as animals.
