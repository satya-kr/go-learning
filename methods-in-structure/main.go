package main

import "fmt"

type Calcu struct {
	num  int
	num2 int
}

func (c Calcu) sum() int {
	return c.num + c.num2
}

func (c Calcu) mul() int {
	return c.num * c.num2
}

func main() {
	cal := Calcu{20, 50}

	fmt.Println("Sum: ", cal.sum())
	fmt.Println("Mul: ", cal.mul())
}
