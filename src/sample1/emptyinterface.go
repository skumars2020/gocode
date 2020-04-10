package main

import "fmt"

func main() {
	Anything("Hello world")
	Anything(22)
	Anything(22.2222)
	Anything(struct{}{}) //Declaring empty struct
}

//Empty declaration of interface we will declare this type of interface when we donot know the type of varable
func Anything(anything interface{}) {
	fmt.Println(anything)
}
