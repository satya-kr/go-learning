package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Course string `json:"course,omitempty"`
}

func main() {
	db := []student{}

	//open file
	file, err := os.Open("studentdb.json")
	if err != nil {
		panic(err)
	}
	//close file end of the program
	defer file.Close()

	//json decode
	json.NewDecoder(file).Decode(&db)
	fmt.Println(db)
}
