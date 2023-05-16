package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter range for Fibonacci Series")
	value := getInput()

	num, _ := strconv.Atoi(value)
	displayFibonacci(num)
}

func getInput() string {
	var data string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data = scanner.Text()
	}
	return data
}

func displayFibonacci(value int) {
	n1 := 0
	n2 := 1
	fmt.Print(n1, ",", n2)
	for i := 0; i < value-2; i++ {
		t := n1 + n2
		n1 = n2
		n2 = t
		fmt.Print(",", t)
	}
	fmt.Println()
}
