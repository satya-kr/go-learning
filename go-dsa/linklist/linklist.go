package main

import (
	"fmt"
	"os"
	"strconv"
)

type Linklist struct {
	number int
	next   *Linklist
}

func (node *Linklist) insert(num int) {
	var temp = &Linklist{}
	temp.number = num
	temp.next = nil
	// fmt.Println(temp)
	// fmt.Println(temp.number)

	if node == nil {
		node = temp
	} else {
		var p = &Linklist{}
		p = node
		for p.next != nil {
			p = p.next
			fmt.Println(p.next)
		}
		p.next = temp
	}
}

func (node *Linklist) display() {
	var p = &Linklist{}
	p = node.next
	for p != nil {
		fmt.Printf("%d -> ", p.number)
		p = p.next
	}
}

func main() {
	fmt.Println("Welcome Mr.Kumar")

	head := new(Linklist)
	var choice string

	for true {
		fmt.Println("Enter the choice")
		fmt.Println("1. Insert value to linklist")
		fmt.Println("2. Display linklist")
		fmt.Println("3. Exit")
		fmt.Scanln(&choice)

		switch choice {

		case "1":
			var data string
			fmt.Printf("Enter value:")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			head.insert(num)
		case "2":
			head.display()
		default:
			os.Exit(0)

		}
	}
}
