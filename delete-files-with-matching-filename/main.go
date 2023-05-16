package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	fmt.Printf("Welcome Mr.Kumar\n\n")

	files, err := ioutil.ReadDir("files")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Enter File name :")
	val := getInput()
	for _, file := range files {
		match, _ := regexp.MatchString(val, file.Name())

		if match {
			filepath, err := filepath.Abs("./files/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			err = os.Remove(filepath)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(file.Name() + " File Deleted")
			}
		} else {
			fmt.Println("File not found!")
		}
	}
}

func getInput() string {
	var data string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data = scanner.Text()
	}
	return data
}
