package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"fmt"
)

func main(){
	//PerformGetRequest()

	// PerformPostJsonRequest()

	PerformPostFormRequest()
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

func PerformPostJsonRequest()  {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename" : "golang",
			"price" : 9,
			"platfom" : "platform.cc"
		}
	`)
	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content , _ := ioutil.ReadAll(res.Body)

	fmt.Println("Post-Response: ", string(content))
	
}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("fname", "clanky")
	data.Add("lname", "bruke")
	data.Add("email", "clanky.bro")

	res, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content , _ := ioutil.ReadAll(res.Body)

	fmt.Println("Post-Form-Response: ", string(content))
}