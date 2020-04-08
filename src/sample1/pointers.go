package main

import "fmt"

type Car struct {
	Model string
	Year  int
	Name  string
}

func main() {
	// Array Example abstract data type
	arr := []int{1, 2, 3}
	arr1 := []string{"Hi", "Siva", "Kumar"}
	fmt.Println(arr1)
	fmt.Println(arr)

	//Pointer Example
	//todo:: Pointer is a Just Varable which stores value of another varabules
	g1 := "Siva"
	p1 := &g1
	fmt.Println(*p1)
	m1, m2 := 2, 4
	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)

	//Structure :: Abstract data type which will groups the logically related data
	//Syntic :: type Car struct
	//It should declare out of method

	c := Car{
		Model: "Tata",
		Year:  2020,
		Name:  "Gold",
	}
	fmt.Println(c)
	fmt.Println(c.getModel())
}

func swap(m1, m2 *int) {
	temp := *m2
	*m2 = *m1
	*m1 = temp

}

func (c Car) draving() {
	fmt.Println("Draving")
}

func (c Car) getModel() string {
	return c.Name
}
