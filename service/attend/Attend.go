package attend

import "github.com/luizcavalieri/go-graphql-starter/model"

// User description.
// swagger:model user
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
	Attendee *model.User `json:"attendee,omitempty"`
	// Lesson attended
	//
	// required: false
	Lesson int `json:"lesson,omitempty"`
}
