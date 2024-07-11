package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func readFile(file string) (string, error) {
	data, err := os.ReadFile(fmt.Sprintf("%v.json", file))
	if err != nil {
		log.Fatal(err)
	}

	return string(data), nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func users(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(vars)
	data, _ := readFile(id)
	fmt.Fprint(w, data)
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/response/{id}", users)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	handleRequests()
}
