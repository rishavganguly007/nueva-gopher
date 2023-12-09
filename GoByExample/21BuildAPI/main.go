package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Moddel for course

type Course struct {
	CourseId string `json:"course"`
	CourseName string `json:"coursename"`
	Price string `json:"price"`
	Author *Author `json:"author"`

}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website`
}

// my fake db

var courses []Course

// middleware, helper
func(c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {

}

// controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from req
	params := mux.Vars(r)

	// loop through courses, find matching id and return resp

	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One course")
	w.Header().Set("Content-Type", "application/json")

	// if: body empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	// if: data -> []
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty(){
		json.NewEncoder(w).Encode("please send some data")
		return
	}

	// generate unique id, string
	// append course ito courses

	course.CourseId = strconv.Itoa(rand.Intn(100))
	json.NewEncoder(w).Encode(course)
	return 
}