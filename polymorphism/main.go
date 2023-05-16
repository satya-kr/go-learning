package main

import "fmt"

type Calculation interface {
	calculate() int
	displayResult()
}

func (s Sum) calculate() int {
	return s.n1 + s.n2
}

func (s Mul) calculate() int {
	return s.n1 * s.n2
}

func (s Sum) displayResult() {
	fmt.Printf("\nSum=%d", s.calculate())
}

func (s Mul) displayResult() {
	fmt.Printf("\nMul=%d", s.calculate())
}

type Sum struct {
	n1 int
	n2 int
}

type Mul struct {
	n1 int
	n2 int
}

func main() {
	fmt.Println("Welcome Mr.Kumar")

	s := Sum{25, 55}
	m := Mul{2, 5}

	cc := []Calculation{s, m}
	for _, c := range cc {
		c.displayResult()
	}
}
