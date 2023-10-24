package main

import (
	"encoding/json"
	// "io/ioutil"
	// "net/http"
	// "net/url"
	// "strings"

	"fmt"
)

type course struct {
	Name string `json:"courseName"` // creating alias with "json" lib
	Price int
	Platform string
	Password string `json:"-"` // "-" -> masks the field
	Tags []string `json:"tags,omitempty"` // omitempty removes the empty fields
}

func main(){

	DecodeJson()
}

func EncodeJson() []byte{
	lcoCourse := [] course{
		{"React", 299, "platform", "pass", []string{"t1", "t2"}},
		{"mern", 299, "platform", "pass", []string{"t1", "t2"}},
		{"mean", 299, "platform", "pass", nil},
	}

	// pkg this data a json data

	jsonData, err := json.MarshalIndent(lcoCourse, "", "\t")


	if err != nil {
		panic(err)
	}
	//fmt.Printf("json : %s/n", jsonData)

	return jsonData
}

func DecodeJson() {
	jsonData := EncodeJson()

	var lcoCourse []course 
	if json.Valid(jsonData) {
		fmt.Printf("valid json \n")

		json.Unmarshal(jsonData, &lcoCourse)
	//	fmt.Printf("%#v\n", lcoCourse)
	} else{
		fmt.Println("Invalid JSON")
	}

// case: we need key-val pair

	var dta []map[string]interface{}
	json.Unmarshal(jsonData, &dta)
	fmt.Printf("%#v\n", dta)

}