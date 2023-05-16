package main

import (
	"fmt"
	"os"
)

func createFile() {
	file, err := os.Create("info.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	data := []byte("here is some information")
	file.Write(data)
	file.WriteString("\nhere is another way to write data into file")
}

func openFile() {
	file, err := os.Open("info.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(data)
	fmt.Println(string(data))
}

func openFile1() {
	data, err := os.ReadFile("info.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func main() {
	// createFile()
	// openFile()
	openFile1()
}
