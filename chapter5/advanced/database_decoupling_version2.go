package advanced

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Data struct {
	Line string
}


type Xenia struct {
	Host string
	Timeout time.Duration
}

func ( x *Xenia) Pull(d *Data) error {
	// Randomly return some data or error

	switch rand.Intn(100) {
	case 1, 10, 90: return io.EOF
	case 5 : return errors.New("data not found ")
	default:
		d.Line = "This is data from Xenia"
	}
	return nil
}

type Pillar struct {
	Host string
	Timeout time.Duration
}

func ( p *Pillar) Store(d *Data) error {
	fmt.Printf("%s",d.Line)
	return nil
}

// Puller is the interface for source system.
type Puller interface {
	Pull(d *Data) error
}

// Storer is the interface for the destination system.
type Storer interface {
	Store(d *Data) error
}

func pull(x Puller, data []Data) (int, error) {
	for i := range data {
		if err := x.Pull(&data[i]);  err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func store(s Storer, data[]Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Copy(p Puller, s Storer, batch int) error {
	data := make([]Data, batch)
	for {

		i,err := pull(p, data)
		if i > 0 {
			i, err = store(s,data[:i])
		}

		if err != nil {
			return err
		}
	}
}

func main() {

	x := Xenia{
		Host:    "localhost:8080",
		Timeout: time.Second,
	}
	p := Pillar{
		Host:    "locahost:9090",
		Timeout: time.Second,
	}


	if err :=Copy(&x , &p, 3	); err != io.EOF {
		fmt.Printf("Ran into errors %s", err)
		return
	}
	fmt.Println("Copy done..")
}