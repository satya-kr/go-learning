package main

import "fmt"

type Calcu struct {
	num  int
	num2 int
}

//interface
type Calculation interface {
	sum() int
	mul() int
}

//methods
func (c Calcu) sum() int {
	return c.num + c.num2
}

func (c Calcu) mul() int {
	return c.num * c.num2
}

//use interface

func useCalculationInf(c Calculation) {
	fmt.Println("Sum:", c.sum())
	fmt.Println("Mul:", c.mul())
}

func main() {
	cal := Calcu{20, 50}

	useCalculationInf(cal)
}
