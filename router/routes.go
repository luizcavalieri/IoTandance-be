package router

import (
	"net/http"

	handler "github.com/luizcavalieri/IoTendance-be/service/users"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"GetUsers",
		"GET",
		"/users",
		handler.GetUsers,
	},
	Route{
		"GetPerson",
		"GET",
		"/users/{id}",
		handler.GetUser,
	},
}
