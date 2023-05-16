package main

import "fmt"

type student struct {
	firstname string
	lastname  string
	bio       string
}

func (s student) fullname() string {
	return fmt.Sprintf("%s %s", s.firstname, s.lastname)
}

type topic struct {
	title   string
	content string
	student //<- here we are doing composition
}

func (t topic) details() {
	fmt.Println("Title:", t.title)
	fmt.Println("Content:", t.content)
	fmt.Println("Student:", t.fullname())
	fmt.Println("Bio:", t.bio)
}

type website struct {
	topics []topic
}

func (w website) contents() {
	fmt.Println("Content of website\n")

	for _, v := range w.topics {
		v.details()
		fmt.Println("")
	}
}

func main() {
	// Golang Composition as a Replacement of Inheritance
	fmt.Println("Welcome Mr.Kumar")

	s1 := student{
		"Akash", "Kumar", "PHP-DEV",
	}
	t1 := topic{
		"Struct insted of class in go", "Go support class but methods can be add to struct", s1,
	}

	t2 := topic{
		"Inheretance in Go", "Go support composition insted of inheratance", student{
			"Mr.Satyajit", "Kumar", "Go Dev",
		},
	}

	t3 := topic{
		"Concurrency in Go", "Go is concurrent language not a parallel one", s1,
	}

	w := website{
		topics: []topic{t1, t2, t3},
	}
	w.contents()
}
