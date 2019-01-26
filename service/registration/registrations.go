package registration

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

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

var registrarsCls []RegistrarClass

// Display a single data
func GetLessonEnrollmentsByUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /attendance/user/{userId}/lesson/{lessonId}
	//
	// Lists attendees enrolled in a lesson.
	//
	// This will show a list with every attendee enrolled in a lesson.
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

	var registrarCls RegistrarClass
	registrarsCls = []RegistrarClass{}

	rows, err := driver.Db.Query(
		"Select st.fname," +
			" 		   st.lname," +
			"		   st.prefname," +
			"		   cl.start_date," +
			"          cl.end_date," +
			"	       cl.class_id, " +
			"		   en.end_date," +
			"          en.commenced," +
			"	       ls.lesson_date," +
			"          ls.lesson_id," +
			"          ts.dayofweek, " +
			"          ts.start_time," +
			"          ts.end_time," +
			"          st.id," +
			"		   rm.name " +
			"from attendee st " +
			"	inner join registration en on en.attendee_id = st.id " +
			" 	inner join class cl on en.class_id = cl.class_id " +
			"	inner join lesson ls  on ls.class_id = cl.class_id " +
			"	inner join rooms rm on rm.room_id = ls.lesson_room " +
			"	inner join timeslots ts on ts.slot_id = ls.lesson_timeslot and " +
			"						 ls.lesson_teacher = " + usrId + " and " +
			" 						 ls.lesson_id = " + lessonId +
			"")
	global.LogFatal(err, "Attendance for user failed")

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&registrarCls.Attendee.FirstName,
			&registrarCls.Attendee.LastName,
			&registrarCls.Attendee.PreferredName,
			&registrarCls.Class.StartDate,
			&registrarCls.Class.EndDate,
			&registrarCls.Class.ID,
			&registrarCls.Registration.EndDate,
			&registrarCls.Registration.Commenced,
			&registrarCls.Lesson.LessonDate,
			&registrarCls.Lesson.ID,
			&registrarCls.TimeSlot.DayOfWeek,
			&registrarCls.TimeSlot.StartTime,
			&registrarCls.TimeSlot.EndTime,
			&registrarCls.Attendee.ID,
			&registrarCls.Room.Name,
		)
		global.LogFatal(err, "Not possible to add record to results.")

		w.WriteHeader(http.StatusOK)
		registrarsCls = append(registrarsCls, registrarCls)

	}
	json.NewEncoder(w).Encode(registrarsCls)

}

// TODO: move to registration service
// Display a single data
func GetLessonEnrollments(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /attendance/lesson/{lessonId}
	//
	// Lists attendees enrolled in a lesson.
	//
	// This will show a list with every attendee enrolled in a lesson.
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

	lessonId := params["lessonId"]
	log.Println("Getting registrations for lesson", lessonId)

	driver.DbInit()

	var registrarCls RegistrarClass
	registrarsCls = []RegistrarClass{}

	rows, err := driver.Db.Query(
		"Select st.fname," +
			" 		   st.lname," +
			"		   st.prefname," +
			"		   cl.start_date," +
			"          cl.end_date," +
			"	       cl.class_id, " +
			"		   en.end_date," +
			"          en.commenced," +
			"	       ls.lesson_date," +
			"          ls.lesson_id," +
			"          ts.dayofweek, " +
			"          ts.start_time," +
			"          ts.end_time," +
			"          st.id," +
			"		   rm.name " +
			"from attendee st " +
			"	inner join registration en on en.attendee_id = st.id " +
			" 	inner join class cl on en.class_id = cl.class_id " +
			"	inner join lesson ls  on ls.class_id = cl.class_id " +
			"	inner join rooms rm on rm.room_id = ls.lesson_room " +
			"	inner join timeslots ts on ts.slot_id = ls.lesson_timeslot and " +
			" 						 ls.lesson_id = " + lessonId +
			"")
	global.LogFatal(err, "Attendance for user failed")

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&registrarCls.Attendee.FirstName,
			&registrarCls.Attendee.LastName,
			&registrarCls.Attendee.PreferredName,
			&registrarCls.Class.StartDate,
			&registrarCls.Class.EndDate,
			&registrarCls.Class.ID,
			&registrarCls.Registration.EndDate,
			&registrarCls.Registration.Commenced,
			&registrarCls.Lesson.LessonDate,
			&registrarCls.Lesson.ID,
			&registrarCls.TimeSlot.DayOfWeek,
			&registrarCls.TimeSlot.StartTime,
			&registrarCls.TimeSlot.EndTime,
			&registrarCls.Attendee.ID,
			&registrarCls.Room.Name,
		)
		global.LogFatal(err, "Not possible to add record to results.")

		w.WriteHeader(http.StatusOK)
		registrarsCls = append(registrarsCls, registrarCls)

	}
	json.NewEncoder(w).Encode(registrarsCls)

}
