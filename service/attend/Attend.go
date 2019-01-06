package attend

import "github.com/luizcavalieri/IoTendance-be/service/users"

// Attend description.
// swagger:model attend
type Attend struct {
	// ID of the attend
	//
	// required: true
	ID int32 `json:"id,omitempty"`
	// How many hours attendee was present
	//
	// required: true
	HoursAttend int8 `json:"hours_attend,omitempty"`
	// Late if attendee is late to the session
	//
	// required: false
	Late int8 `json:"late,omitempty"`
	// Attendee
	//
	// required: false
	Attendee *Attendee `json:"attendee,omitempty"`
	// Lesson attended
	//
	// required: false
	Lesson int `json:"lesson,omitempty"`
}

// Attendee description.
// swagger:model attendee
type Attendee struct {
	// ID of the attend
	//
	// required: true
	ID int32 `json:"id,omitempty"`
	// Attendee first name
	//
	// required: true
	FirstName int8 `json:"fname,omitempty"`
	// Attendee last name
	//
	// required: true
	LastName int8 `json:"lname,omitempty"`
	// Attendee preferred name
	//
	// required: true
	PreferredName int8 `json:"prefname,omitempty"`
	// Attendee start date
	//
	// required: true
	StartDate string `json:"enrol_date,omitempty"`
	// User entity
	//
	// required: true
	User *users.User `json:"attendee,omitempty"`
}
