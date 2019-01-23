package lesson

// Lesson description.
// swagger:model lesson
type Lesson struct {
	// ID of the lesson
	//
	// required: true
	ID int `json:"lesson_id,omitempty"`
	// Class id of the class
	//
	// required: true
	Class string `json:"class_id,omitempty"`
	// Lesson number
	//
	// required: true
	LessonNo int64 `json:"lesson_no,omitempty"`
	// Lesson date
	//
	// required: true
	LessonDate string `json:"lesson_date,omitempty"`
	// Lesson hours of duration
	//
	// required: false
	LessonHours int `json:"lesson_hours,omitempty"`
	// Lesson's time slot allocated
	//
	// required: true
	TimeSlot int `json:"lesson_timeslot,omitempty"`
	// Lesson's room
	//
	// required: true
	Room int `json:"lesson_room,omitempty"`
	// Lesson's teacher
	//
	// required: true
	Teacher int `json:"lesson_teacher,omitempty"`
	// Lesson's assistant
	//
	// required: false
	LessonAssistant int `json:"lesson_assistant,omitempty"`
}
