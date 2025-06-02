package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/object1", object1Handler).Methods("GET")
	r.HandleFunc("/api/object2", object2Handler).Methods("GET")

	fmt.Println("Listening on :8080…")
	http.ListenAndServe(":8080", r)
}
