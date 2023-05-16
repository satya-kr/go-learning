package main

import (
	"errors"
	"fmt"
)

func checkMySQL() bool {
	return false
}

func checkDBConnection() error {
	if checkMySQL() {
		return nil
	} else {
		return errors.New("MySQL is not running....")
	}
}

func main() {
	if err := checkDBConnection(); err != nil {
		// fmt.Println(err)
		// panic(err.Error())
		defer func() { //check panic is run or not
			if r := recover(); r != nil {
				fmt.Printf("\nPanic Error: %s\n", r)
				fmt.Println("Backup server is running...")
				fmt.Println("My work is going on backup server...")
			}
		}()
		panic("MySQL is not running, run it first then start your program...")
	} else {
		fmt.Println("MySQL is running...")
	}
	fmt.Println("You can do your transaction in MySQL ...")
}
