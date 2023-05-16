package main

import (
	student "encap/student"
	"fmt"
)

func main() {
	fmt.Println("Welcome Mr.Kumar")

	s := student.Student{Rollno: 1}
	// s.Rollno = 5
	s.SetName("Abhijit Kumar [Kata Paa]")
	s.SetCourse("MCA")

	fmt.Println("Student Name : ", s.GetName())
	fmt.Println("Student Course : ", s.GetCourse())
	fmt.Println("Student Rollno : ", s.GetRollno())

	s1 := student.Student{Rollno: 2}
	s1.Rollno = 3
	s1.SetName("Satyajit Kumar [DADA - Khauf Ka Dusra Name]")
	s1.SetCourse("BCA")

	fmt.Println("\nStudent Name : ", s1.GetName())
	fmt.Println("Student Course : ", s1.GetCourse())
	fmt.Println("Student Rollno : ", s1.GetRollno())
}
