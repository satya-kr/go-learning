package main

import (
	"fmt"
	"net/http"
	"sync"
)

var reStoreEndpoints = []string{"test-endpoint"}
var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	// go greeter2("Satyajit")
	// go greeter("Akash")
	// greeter("Kumar")

	websites := []string{
		"https://google.com",
		"https://astergo.in",
		"https://bnmotors.in",
		"https://cloudsentinel.tk",
		"https://fb.com",
		"https://github.com",
	}

	for _, url := range websites {
		go getStatusCode(url)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(reStoreEndpoints)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(30 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func greeter2(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Opps erro found in endoint")
	} else {
		//so here we lock the memory untill our endpoints will store
		mut.Lock()
		reStoreEndpoints = append(reStoreEndpoints, endpoint)
		mut.Unlock()
		//so once all of its done unloack the memory
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
	defer wg.Done()
}
