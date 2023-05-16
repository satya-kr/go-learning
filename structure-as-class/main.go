package main

import (
	"fmt"
	student "opps/student"
)

func main() {
	fmt.Println("Welcome Mr.Kumar")

	//obj of student
	// s := student.Student{Name: "Satyajit Kumar", Course: "BCA", Rollno: 1}
	s := student.New("Satyajit Kr", "GO lang", 13)
	s.GetStudentInfo()
	var s1 = student.New("Akash Kr", "PHP & MySQL", 13)
	s1.GetStudentInfo()

}
