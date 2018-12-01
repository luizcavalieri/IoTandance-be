package main

import (
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"

	ppl "github.com/luizcavalieri/IoTendance-be/service/people"
)

func init() {
	gotenv.Load()
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/people", ppl.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", ppl.GetUser).Methods("GET")
	router.HandleFunc("/people/{id}", ppl.CreateUser).Methods("POST")
	router.HandleFunc("/people/{id}", ppl.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8085", router))
}
