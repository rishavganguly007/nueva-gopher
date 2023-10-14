package main

import (
	// "io/ioutil"
	// "net/http"
	"net/url"

	"fmt"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&payments=bfhgdii5g"

func main() {
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("The type of query params: ", qparams)
}