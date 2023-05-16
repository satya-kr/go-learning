package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Printf("Welcome Mr.Kumar\n\n")

	if len(os.Args) < 2 {
		log.Fatalln("Use like: go run main.go url_1 url_2 url_3 ..... url_n")
	}
	// fmt.Println(os.Args[1:])

	for _, url := range os.Args[1:] {
		wg.Add(1)
		checkUrlConn("https://" + url)
	}
	wg.Wait()
}

func checkUrlConn(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status Code : [%d] of URL : %s\n", res.StatusCode, url)
}
