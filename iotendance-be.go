package iotendancebe

import (
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"

	usr "github.com/luizcavalieri/IoTendance-be/service/user"
)

func init() {
	gotenv.Load()
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", usr.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", usr.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", usr.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", usr.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8085", router))
}

