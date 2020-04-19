package main

import (
	"fmt"
	"time"
)

//Goroutines --> Goroutines is not a funcation but Goroutines is a way we call funcation
//it means funcation will going to run in the background
//#Asyc and Concorancy type
func main() {
	go heavy()
	fmt.Println("=========Fin=========")
	go superHeavy()
	select {}
}
func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}
func superHeavy() {
	for {
		time.Sleep(time.Second * 3)
		fmt.Println("SuperHeavy")
	}
}
