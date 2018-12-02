package main

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/subosito/gotenv"

	"github.com/luizcavalieri/IoTendance-be/global"
)


func init() {
	gotenv.Load()
}

// main function to boot up everything
func main() {

	s := `
                schema {
                        query: Query
                }
                type Query {
                        hello: String!
                }
        `
	schema := graphql.MustParseSchema(s, &global.Query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8085", nil))


	//router := mux.NewRouter()
	//router.HandleFunc("/people", usr.GetUsers).Methods("GET")
	//router.HandleFunc("/people/{id}", usr.GetUser).Methods("GET")
	//router.HandleFunc("/people/{id}", usr.CreateUser).Methods("POST")
	//router.HandleFunc("/people/{id}", usr.DeleteUser).Methods("DELETE")
	//
	//log.Fatal(http.ListenAndServe(":8085", router))
}
