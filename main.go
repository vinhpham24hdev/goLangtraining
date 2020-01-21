package main

import (
	control "goClean/controller"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/rates", control.GetRates).Methods("GET")
	r.HandleFunc("/rates/{day}", control.GetRate).Methods("GET")
	r.HandleFunc("/avgrates", control.GetAvgRate).Methods("GET")
	log.Println("Listen localhost:8000")
	http.ListenAndServe(":8000", r)
}
