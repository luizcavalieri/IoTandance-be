package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"log"
	ppl "github.com/luizcavalieri/iotendance-be/service/people"
)


var people []ppl.Person

// main function to boot up everything
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/people", ppl.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", ppl.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", ppl.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", ppl.DeletePerson).Methods("DELETE")

	ppl.InitPeople()
	log.Fatal(http.ListenAndServe(":8080", router))
}

