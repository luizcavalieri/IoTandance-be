package registration

import (
	"github.com/luizcavalieri/iotendance-be/service/attend"
	"github.com/luizcavalieri/iotendance-be/service/class"
	"github.com/luizcavalieri/iotendance-be/service/lesson"
	"github.com/luizcavalieri/iotendance-be/service/room"
	"github.com/luizcavalieri/iotendance-be/service/timeslot"
	"time"
)

// Registration description.
// swagger:model registration
type Registration struct {
	// ID of the lesson
	//
	// required: true
	ID string `json:"registration_id,omitempty"`
	// Class id
	//
	// required: true
	Class string `json:"class_id,omitempty"`
	// Attendee is
	//
	// required: true
	Attendee int `json:"attendee_id,omitempty"`
	// Attendee commenced
	//
	// required: true
	Commenced int8 `json:"commenced,omitempty"`
	// Start date of enrollment
	//
	// required: true
	StartDate string `json:"start_date,omitempty"`
	// End date of enrollment
	//
	// required: true
	EndDate time.Time `json:"end_date,omitempty"`
	// Results of enrollment
	//
	// required: false
	Result int `json:"result,omitempty"`
	// Comments on enrollment
	//
	// required: false
	Comment string `json:"comment,omitempty"`
	// Final date of enrollment
	//
	// required: false
	FinalDate string `json:"final_date,omitempty"`
	// Qualification
	//
	// required: false
	Qualification int `json:"qualification,omitempty"`
}

// RegistrationClass description.
// swagger:model AttendanceClass
type RegistrarClass struct {
	// Attendee of the class
	//
	// required: true
	Attendee attend.Attendee `json:"attendee"`
	// Class first name
	//
	// required: true
	Class class.Class `json:"class"`
	// Registration last name
	//
	// required: true
	Registration Registration `json:"registration"`
	// Lesson preferred name
	//
	// required: true
	Lesson lesson.Lesson `json:"lesson,omitempty"`
	// TimeSlot start date
	//
	// required: true
	TimeSlot timeslot.TimeSlot `json:"timeslot"`
	// Room entity
	//
	// required: true
	Room room.Room `json:"room"`
	// Attend entity
	//
	// required: false
	Attend attend.Attend `json:"attend,omitempty"`
}
