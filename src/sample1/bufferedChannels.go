package main

import "fmt"

type car struct {
	Model string
}

func main() {
	c := make(chan *car, 3)

	go func() {
		c <- &car{"1"}
		c <- &car{"2"}
		c <- &car{"3"}
		c <- &car{"4"}
		close(c)
	}()

	for r := range c {
		fmt.Println(r.Model)

	}

}
