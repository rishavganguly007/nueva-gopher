package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	fmt.Println("Backend with Golang")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "reactJS", Price: "299", Author: &Author{Fullname: "Rg",
								Website: "abc.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "JAVA", Price: "299", Author: &Author{Fullname: "Rg",
	Website: "abc.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourses).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

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

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// lopp, id, remove, add with given id

	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index + 1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop,id, remove

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}