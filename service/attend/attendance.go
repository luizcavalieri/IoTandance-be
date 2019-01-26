package attend

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"github.com/luizcavalieri/IoTendance-be/driver"
	"github.com/luizcavalieri/IoTendance-be/global"
)

// IDParam is used to identify a user
//
// swagger:parameters listUser
type IDParam struct {
	// The ID of a user
	//
	// in: path
	// required: true
	ID string `json:"user_id"`
}

var attends []Attend

func GetAttendance(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /people people listPeople
	//
	// Lists all users.
	//
	// This will show all recorded people.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: usersResponse

	log.Println("Get attendance")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	driver.DbInit()
	var attend Attend
	attends = []Attend{}

	rows, err := driver.Db.Query("SELECT * from attend")
	global.LogFatal(err, "")

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&attend.ID, &attend.Attendee, &attend.Lesson, &attend.HoursAttend, &attend.Late)
		global.LogFatal(err, "")

		attends = append(attends, attend)
	}

	json.NewEncoder(w).Encode(attends)
}

// Display a single data
func GetAttendeeLessonAttendance(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /attendance/attendee/{attendeeId}/lesson/{lessonId} Attendance
	//
	// Lists attendance for the attendee in the lesson.
	//
	// This will show the record of attendance in a lesson for an attendee.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Params:
	//     - id: IDParam
	//
	//     Responses:
	//       200: userResponse
	//       404: jsonError

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)

	attendeeId := params["attendeeId"]
	lessonId := params["lessonId"]
	log.Println("Getting attendance for user in lesson", attendeeId, lessonId)

	driver.DbInit()
	var attend Attend
	attends = []Attend{}

	lessonIdInt, err := strconv.ParseInt(params["lessonId"], 10, 64)
	if err != nil {
		global.LogFatal(err, "Parsing attendeeId string to int")
	}

	attendeeIdInt, err := strconv.ParseInt(params["attendeeId"], 10, 64)
	if err != nil {
		global.LogFatal(err, "Parsing attendeeId string to int")
	}
	rows, err := driver.Db.Query(
		"" +
			"SELECT * " +
			"	from attend" +
			"   where attendee_id =" + attendeeId +
			"     and lesson_id = " + lessonId +
			"")
	global.LogFatal(err, "Attendance for user failed")

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&attend.ID,
			&attend.Attendee,
			&attend.Lesson,
			&attend.HoursAttend,
			&attend.Late,
		)

		if attend.Attendee == attendeeIdInt && attend.Lesson == lessonIdInt {
			w.WriteHeader(http.StatusOK)
			attends = append(attends, attend)
		}

		global.LogFatal(err, "Not possible to add record to results.")

	}
	json.NewEncoder(w).Encode(attends)

}
