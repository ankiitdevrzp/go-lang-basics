package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"} //pointer
var mut sync.Mutex             //pointer

var wg sync.WaitGroup

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://github.com",
		"https://fb.com",
		"https://youtube.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Microsecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code fro %s\n", res.StatusCode, endpoint)
	}
}
