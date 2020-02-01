package chapter5

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct{}

func (bike) Move() {
	fmt.Println("Bike can move.")
}

func (bike) Lock() {
	fmt.Println("Bike can lock.")
}

func (bike) Unlock() {
	fmt.Println("Bike an unlock.")
}

func main() {
	var ml MoveLocker
	var m Mover

	ml = bike{}
	m = ml // This is Ok.

	// ml = m // This is no Ok.
	b, ok := m.(bike) // Extract the underlying concrete type and then assign it to the MoveLocker interface.
	if ok {
		ml = b
	}
}