package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/subosito/gotenv"

	"github.com/luizcavalieri/IoTendance-be/router"
)

func init() {
	gotenv.Load()
}


// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// main function to boot up everything
func main() {

	// create router and start listen on port 8000
	route := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8002", setupGlobalMiddleware(route)))
	//
	//router.HandleFunc("/people", ppl.GetPeople).Methods("GET")
	//router.HandleFunc("/people/{id}", ppl.GetUser).Methods("GET")
	//router.HandleFunc("/people/{id}", ppl.CreateUser).Methods("POST")
	//router.HandleFunc("/people/{id}", ppl.DeleteUser).Methods("DELETE")

	//log.Fatal(http.ListenAndServe(":8085", router))
}
