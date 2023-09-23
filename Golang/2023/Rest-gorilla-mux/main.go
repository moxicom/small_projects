package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world")
	router := mux.NewRouter()
	//	router.StrictSlash(false)
	router.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Tasks post received")
	}).Methods("POST")

	router.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(mux.Vars(r))
	}).Methods("GET")

	router.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {}).Methods("DELETE")
	router.HandleFunc("/tasks/id:[0-9]+", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	router.HandleFunc("/tags/{tag}", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	router.HandleFunc("/due/{year:[0-9]+}/{month:[0-9]+}/{day:[0-9]+}/", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	fmt.Println("SERVER IS RUNNING ON localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
