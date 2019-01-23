package class

import "time"

// Class description.
// swagger:model class
type Class struct {
	// ID of the class
	//
	// required: true
	ID string `json:"class_id,omitempty"`
	// Start date of the class
	//
	// required: true
	StartDate string `json:"start_date,omitempty"`
	// Course id number
	//
	// required: true
	Course int64 `json:"course_id,omitempty"`
	// End date of class
	//
	// required: false
	EndDate time.Time `json:"end_date,omitempty"`
	// Days of week class will be held
	//
	// required: false
	DaysOfWeek []int `json:"lesson_hours,omitempty"`
	// Hours per day
	//
	// required: true
	HoursPerDay []int `json:"hours_per_day,omitempty"`
	// Semester
	//
	// required: true
	Semester int `json:"semester_id,omitempty"`
	// Lesson's teacher
	//
	// required: true
	Teacher int `json:"lesson_teacher,omitempty"`
}
