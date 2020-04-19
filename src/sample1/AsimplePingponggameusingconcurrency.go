package main

import (
	"fmt"
	"time"
)

func main() {

	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- 1

	select {}

}

// getting value from chan to pinger in this method and set value
// to poner and sending via chan from the same method
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

// getting value from chan to poner in this method and set value to
// pinger and sending via chan from the same method
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}
