package timeslot

// Lesson description.
// swagger:model lesson
type TimeSlot struct {
	// ID of the time slot
	//
	// required: true
	ID string `json:"slot_id,omitempty"`
	// Day of the week
	//
	// required: true
	DayOfWeek int `json:"dayofweek,omitempty"`
	// Name of time slot
	//
	// required: true
	Name string `json:"name,omitempty"`
	// Start time of time slot
	//
	// required: true
	StartTime string `json:"start_time,omitempty"`
	// End time of time slot
	//
	// required: true
	EndTime string `json:"end_time,omitempty"`
	// Sector of time slot
	//
	// required: true
	Sector int `json:"sector_id,omitempty"`
}
