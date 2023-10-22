package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"fmt"
)

func main(){
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	var resString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	
	byteCount, _  := resString.Write(content)
	fmt.Println("byteCount : ", byteCount)
	fmt.Println(resString.String())
	

	//fmt.Println(string(content))
	
}