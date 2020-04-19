package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch := make(chan int, 2)
	quits := make(chan struct{})
	go func() {
		ch <- 1
		quits <- struct{}{}
	}()

	go Selects(ch, quits)

	select {}

}
func Selects(ch chan int, quits chan struct{}) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ch:
			fmt.Println("got the response")
		case <-quits:
			fmt.Println("Quits")
			os.Exit(0)
		}

	}
}
