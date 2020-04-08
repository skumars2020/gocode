package main

import (
	"fmt"
	"strings"
)

func main() {
	firstname := "Siva Kumar"
	lastName := "Sajja"
	fmt.Println(firstname + " " + lastName)
	fmt.Println(strings.Contains(firstname, lastName))
}
