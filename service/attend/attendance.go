package attend

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"github.com/luizcavalieri/IoTendance-be/driver"
	"github.com/luizcavalieri/IoTendance-be/global"
	"github.com/luizcavalieri/IoTendance-be/service/class"
	"github.com/luizcavalieri/IoTendance-be/service/lesson"
	"github.com/luizcavalieri/IoTendance-be/service/registration"
	"github.com/luizcavalieri/IoTendance-be/service/room"
	"github.com/luizcavalieri/IoTendance-be/service/timeslot"
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
	// swagger:route GET /users/{id} users listUsers
	//
	// Lists user from their id.
	//
	// This will show the record of an identified user.
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
		"SELECT * from attend where attendee_id =" + attendeeId + " and lesson_id = " + lessonId)
	global.LogFatal(err, "Attendance for user failed")

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&attend.ID, &attend.Attendee, &attend.Lesson,
			&attend.HoursAttend, &attend.Late)

		if attend.Attendee == attendeeIdInt && attend.Lesson == lessonIdInt {
			w.WriteHeader(http.StatusOK)
			attends = append(attends, attend)

		}
		global.LogFatal(err, "Not possible to add record to results.")

	}
	json.NewEncoder(w).Encode(attends)

}

// Display a single data
func GetLessonAttendance(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /users/{id} users listUsers
	//
	// Lists user from their id.
	//
	// This will show the record of an identified user.
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

	usrId := params["userId"]
	lessonId := params["lessonId"]
	log.Println("Getting attendance for user in lesson", usrId, lessonId)

	driver.DbInit()
	var attend Attend
	var registr registration.Registration
	var attendee Attendee
	var cls class.Class
	var ls lesson.Lesson
	var timeSlot timeslot.TimeSlot
	var rm room.Room
	attends = []Attend{}

	//lessonIdInt, err := strconv.Atoi(lessonId)
	//if err != nil {
	//	global.LogFatal(err, "Parsing attendeeId string to int")
	//}

	//usrIdInt, err := strconv.Atoi(usrId)
	//if err != nil {
	//	global.LogFatal(err, "Parsing attendeeId string to int")
	//}
	rows, err := driver.Db.Query(
		"Select st.fname, st.lname, st.prefname, cl.start_date, cl.end_date, cl.class_id, " +
			"en.end_date, en.commenced, ls.lesson_date, ls.lesson_id, ts.dayofweek, " +
			"ts.start_time, ts.end_time, st.id, rm.name " +
			"from attendee st " +
			"inner join registration en " +
			"	on en.attendee_id = st.id " +
			"inner join class cl " +
			"	on en.class_id = cl.class_id " +
			"inner join lesson ls " +
			"	on ls.class_id = cl.class_id " +
			"inner join rooms rm " +
			"	on rm.room_id = ls.lesson_room " +
			"inner join timeslots ts " +
			"	on ts.slot_id = ls.lesson_timeslot and " +
			"ls.lesson_teacher = " + usrId + " and " +
			"ls.lesson_id = " + lessonId +
			"")
	global.LogFatal(err, "Attendance for user failed")

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&attendee.FirstName, &attendee.LastName, &attendee.PreferredName,
			&cls.StartDate, &cls.EndDate, &cls.ID, &registr.EndDate, &registr.Commenced, &ls.LessonDate,
			&ls.ID, &timeSlot.DayOfWeek, &timeSlot.StartTime, &timeSlot.EndTime, &attendee.ID, &rm.Name)

		attends = append(attends, attend)

		if err != nil {
			global.LogFatal(err, "Not possible to add record to results.")
		}

	}
	json.NewEncoder(w).Encode(attends)

}
