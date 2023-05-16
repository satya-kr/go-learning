package student

import "fmt"

type student struct {
	Name   string
	Course string
	Rollno int
}

func New(name, course string, rollno int) student {
	s := student{Name: name, Course: course, Rollno: rollno}
	return s
}

func (s *student) GetStudentInfo() {
	fmt.Printf("\n\nName = %s\nCourse = %s\nRollno = %d", s.Name, s.Course, s.Rollno)
}
