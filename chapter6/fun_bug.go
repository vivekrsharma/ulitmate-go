package chapter6

import "fmt"

type errorMessage struct {
	message string
}

func (e *errorMessage) Error() string {
	return e.message
}

func New() error {
	return &errorMessage{"Find the bug"}
}

func returnError() *errorMessage {
	return  nil
}

// The function call to the returnError panics in the line below.
// Can you figure out why? :=)
func main() {
	var err error
	if err = returnError(); err!=nil {
		fmt.Println("How is this possible!")
		panic(err)
	}
	fmt.Println("Lief is good!")

}
