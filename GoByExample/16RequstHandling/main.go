package main

import (
	"io/ioutil"
	"net/http"

	"fmt"
)

const url = "https://lco.dev"
func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(nil)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to clos the conn

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("Content: \n", content)
}