package chapter6

import "fmt"

type customError struct {
	message string
}

func (c customError) Error() string {
	return c.message
}

// New accepts a custom message and returns a new customError.
//noinspection ALL
func New(str string) error {
	return customError{message: str}
}

// Following is the right way to create custom errors.
//func (c *customError) Error() string {
//	return c.message
//}
//
//func New(str string) error {
//	return &customError{message:str}
//}

func webCall() error {
	return New("Bad Request")
}

func main() {
	fakeError := New("Bad Request")
	if err := webCall(); err != nil {
		fmt.Println("This is a webCall error.")
		if fakeError == err {
			fmt.Println("Insanity.. ")
			panic(err)
		}
		fmt.Println("Life is good.")
	}
	fmt.Println("Life is amazing.")

}
