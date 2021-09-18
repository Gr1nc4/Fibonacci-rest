package main

import (
	"log"
	"net/http"

	"github.com/Gr1nc4/Fibonacci-rest.git/fiboService"
	"github.com/gorilla/mux"
)

func HandleFunc() {
	r := mux.NewRouter()
	r.HandleFunc("/", fiboService.GetDiscription).Methods("GET")
	r.HandleFunc("/fibo", fiboService.GetFiboSlice).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))

}
func main() {
	HandleFunc()
}
