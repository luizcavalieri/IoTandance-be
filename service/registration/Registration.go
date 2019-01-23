package registration

import "time"

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
