package main

import (
	// "io/ioutil"
	// "net/http"
	// "net/url"
	// "strings"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){

	fmt.Println("helo")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)


	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter(){
	fmt.Println("Hi")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to Golang</h1>"))
}
