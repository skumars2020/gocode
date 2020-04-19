package main

import (
	"fmt"
	"time"
)

// Channels := are interprocess communications Asyc and syc communications
//Channels help us for iterprocess communictions between Producer and consumer
//Channels it not helps with shared memory
//buffer and unbuffer Channels

func main() {
	fmt.Println("un buffered Channel")
	ch := make(chan int)
	//<Name> chan <datatype>

	go func() {
		ch <- 1
	}()

	p := <-ch
	fmt.Println(p)

	go func() {
		ch <- 2
	}()
	time.Sleep(time.Second * 2)
	x := <-ch
	fmt.Println(x)
	fmt.Println(ch)
}
