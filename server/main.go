package main

import (
	"github.com/luizcavalieri/IoTendance-be/driver"
	"github.com/subosito/gotenv"
	"github.com/vektah/gqlgen-tutorials/dataloader"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"

	iotbe "github.com/luizcavalieri/IoTendance-be/gqlgen"
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

	queryHandler := handler.GraphQL(dataloader.MakeExecutableSchema(dataloader.New(db)))

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(iotbe.NewExecutableSchema(iotbe.Config{Resolvers: &iotbe.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
