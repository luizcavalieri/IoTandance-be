package attend

import (
	"database/sql"

	"github.com/luizcavalieri/iotendance-be/service/user"
)

// Attend description.
// swagger:model attend
type Attend struct {
	// ID of the attend
	//
	// required: true
	ID sql.NullInt64 `json:"id,omitempty"`
	// How many hours attendee was present
	//
	// required: true
	HoursAttend sql.NullInt64 `json:"hours_attend,omitempty"`
	// Late if attendee is late to the session
	//
	// required: false
	Late sql.NullBool `json:"late,omitempty"`
	// Attendee
	//
	// required: false
	Attendee sql.NullInt64 `json:"attendee_id,omitempty"`
	// Lesson attended
	//
	// required: false
	Lesson sql.NullInt64 `json:"lesson,omitempty"`
}

// Attendee description.
// swagger:model attendee
type Attendee struct {
	// ID of the attend
	//
	// required: true
	ID int64 `json:"id,omitempty"`
	// Attendee first name
	//
	// required: true
	FirstName string `json:"fname,omitempty"`
	// Attendee last name
	//
	// required: true
	LastName string `json:"lname,omitempty"`
	// Attendee preferred name
	//
	// required: true
	PreferredName string `json:"prefname,omitempty"`
	// Attendee start date
	//
	// required: true
	StartDate string `json:"enrol_date,omitempty"`
	// User entity
	//
	// required: true
	User *user.User `json:"user,omitempty"`
}
