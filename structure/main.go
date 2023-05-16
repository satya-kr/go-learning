package main

import "fmt"

type Person struct {
	name string
	role string
}

func getPersion() Person {
	p := Person{"Akash Kumar", "Golang Dev"}
	return p
}

func main() {
	p := Person{"Satyajit Kumar", "Web Developer"}
	fmt.Println(p)
	fmt.Println("Name : ", p.name)
	fmt.Println("Job Role : ", p.role)

	p2 := getPersion()
	fmt.Println(p2.name)
	p2.name = "Amit"
	fmt.Println(p2.name)

}
