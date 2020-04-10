package main

import "fmt"

//interface is named collections of method signature
type Car interface {
	Drive()
	Stop()
}

// func NewModel(arg string) Car {
// 	return &Lambo{arg}
// }

type Lambo struct {
	LamboModel string
}

type Tata struct {
	TataModel string
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo is on Move")
	fmt.Println(l.LamboModel)
}
func (t *Tata) Drive() {
	fmt.Println("Tata is on Move")
	fmt.Println(t.TataModel)
}

func main() {
	l := Lambo{"Gallardo"}
	t := Tata{"C396"}
	l.Drive()
	t.Drive()
}
