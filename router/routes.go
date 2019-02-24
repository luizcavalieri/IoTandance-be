package router

import (
	"github.com/luizcavalieri/iotendance-be/service/registration"
	"net/http"

	attendance "github.com/luizcavalieri/iotendance-be/service/attend"
	users "github.com/luizcavalieri/iotendance-be/service/user"
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

	/**** USER ****/
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
		"Login",
		"GET",
		"/login",
		users.Login,
	},

	/**** ATTENDANCE ****/
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

	/**** REGISTRATION ****/
	Route{
		"GetLessonEnrollmentsByUser",
		"GET",
		"/registration/user/{userId}/lesson/{lessonId}",
		registration.GetLessonEnrollmentsByUser,
	},
	Route{
		"GetLessonEnrollments",
		"GET",
		"/registration/lesson/{lessonId}",
		registration.GetLessonEnrollments,
	},
}
