package router

import (
	"net/http"

	attendance "github.com/luizcavalieri/IoTendance-be/service/attend"
	users "github.com/luizcavalieri/IoTendance-be/service/user"
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
		"CreateUser",
		"POST",
		"/users",
		users.CreateUser,
	},
	Route{
		"GetUsers",
		"GET",
		"/users",
		users.GetUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/users/{id}",
		users.GetUser,
	},
	Route{
		"GetAttendance",
		"GET",
		"/attendance",
		attendance.GetAttendance,
	},
	Route{
		"GetAttendeeLessonAttendance",
		"GET",
		"/attendance/attendee/{attendeeId}/lesson/{lessonId}",
		attendance.GetAttendeeLessonAttendance,
	},
	Route{
		"GetLessonAttendance",
		"GET",
		"/attendance/user/{userId}/lesson/{lessonId}",
		attendance.GetLessonAttendance,
	},
}
