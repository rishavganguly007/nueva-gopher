package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // pointer
var mut sync.Mutex // pointer

func main(){
	// go greeter("hellow") // creates another thread
	// greeter("world")

	websites := []string{
		"https://go.dev",
		"https://google.com",
	}

	for _, web := range websites{
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() // holds the the main method from finishing, until it receives signal from wg.Done()
	fmt.Println(signals)
}

// func greeter(s string){
// 	for i:= 0; i <5 ; i++{
// 	fmt.Println(s)
// 	} }

func getStatusCode(endpoint string){
	defer wg.Done() // wait-group has the responsibilty of signalling that this method is done and signals main to finish

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("err in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint )
	}

}