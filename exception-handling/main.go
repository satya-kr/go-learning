package main

import (
	"fmt"
)

type dividedbyzero struct {
	textError string
}

// overriding per-defined Error() method
func (d dividedbyzero) Error() string {
	return d.textError
}

func errorDevidedByZero() error {
	return dividedbyzero{"Divided by zero using structure"}
}

func division(n1, n2 int) (int, error) {
	if n2 == 0 {
		// return -1, errors.New("Divide by zero")
		return 1, errorDevidedByZero()
	}
	return n1 / n2, nil
}

func main() {
	res, err := division(20, 0)
	if err != nil {
		// fmt.Println(err)
		fmt.Println(err.Error())
	} else {
		fmt.Println("Divisdion of two numbers: ", res)
	}
}
