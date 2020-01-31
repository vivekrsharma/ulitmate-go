package chapter5

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

type System struct  {
	Xenia
	Pillar
}

func pull(x *Xenia, data []Data) (int, error) {
	for i := range data {
		if err := x.Pull(&data[i]);  err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func store(p *Pillar, data[]Data) (int, error) {
	for i := range data {
		if err := p.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func Copy(s *System, batch int) error {
	data := make([]Data, batch)
	for {

		i,err := pull(&s.Xenia, data)
		if i > 0 {
			 i, err = store(&s.Pillar,data[:i])
		}

		if err != nil {
			return err
		}
	}
}

func main() {
	s := System {
		Xenia{
			Host:    "localhost:8080",
			Timeout: time.Second,
		},
		Pillar{
			Host:    "localhost:8080",
			Timeout: time.Second,
		},
	}
	if err := Copy(&s, 3); err != io.EOF {
		fmt.Printf("Ran into errors %s", err)
		return
	}
	fmt.Println("Copy done..")
}