package main

import "fmt"

func main() {
	//variable is a name of memory location that holds value at runtime

	var num int = 20
	fmt.Println(num)
	fmt.Println(&num)

	var ptr_num = &num
	fmt.Println(ptr_num)
	fmt.Println(*ptr_num)

	//create another pointer
	var p *int
	var num2 = 45
	p = &num2
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(num2)

	myNum := 24
	var ptr = &myNum
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 55
	fmt.Println(myNum)

	n1 := 20
	n2 := 30
	fmt.Println(n1, n2)
	swap(&n1, &n2)
	fmt.Println(n1, n2)
}

func swap(n1, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}
