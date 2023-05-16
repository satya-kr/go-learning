package main

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"strconv"
)

type student struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Course string `json:"course,omitempty"`
}

func getRandID() string {
	return strconv.Itoa(rand.Intn(100))
}

func main() {

	db := []student{
		{getRandID(), "Satyajit", "Laravel"},
		{getRandID(), "Abhijit", "React JS"},
		{getRandID(), "Akash", "Go Lang"},
	}

	buff := new(bytes.Buffer)
	json.NewEncoder(buff).Encode(db)

	//create file
	file, err := os.Create("studentdb.json")
	if err != nil {
		panic(err)
	}
	//close file end of the program
	defer file.Close()

	io.Copy(file, buff)
}
