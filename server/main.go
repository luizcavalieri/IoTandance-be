package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"

	"github.com/luizcavalieri/IoTendance-be/driver"
	iotendancebe "github.com/luizcavalieri/IoTendance-be/gqlgen"
	"github.com/subosito/gotenv"
)

const defaultPort = "8085"

func init(){
	gotenv.Load()
}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	driver.DbInit()

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(iotendancebe.NewExecutableSchema(iotendancebe.Config{Resolvers: &iotendancebe.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
