package main

import "fmt"

func main() {
	s := sum(10, 25)
	fmt.Println("Sum of numbers: ", s)
	m := mul(10, 25)
	fmt.Println("Multiplication of numbers: ", m)

	sum, mul := Calcu(20, 30)

	fmt.Println("Sum and Multiplication of numbers: ", sum, mul)

	res := varadic_func(1, 52, 65, 32, 14, 58, 23, 59, 64)
	fmt.Println("Sum using varadic_func: ", res)

	sum2 := func(n1, n2 int) int {
		return n1 + n2
	}

	fmt.Println("sum of two numbers", sum2) //this function not like function its as an object
	fmt.Println("sum of two numbers", sum2(20, 50))
}

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func mul(n1, n2 int) int {
	return n1 * n2
}

//return multiple values
func Calcu(n1, n2 int) (int, int) {
	s := n1 + n2
	m := n1 * n2

	return s, m
}

//varadic function pass multiple arguments
func varadic_func(nums ...int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}
