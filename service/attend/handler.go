package attend

// JsonError is a generic error in JSON format
//
// swagger:response jsonError
type jsonError struct {
	// in: body
	Message string `json:"message"`
}

// AttendanceResponse contains a single person information
//
// swagger:response attendResponse
type attendanceResponse struct {
	// in: body
	Payload *Attend `json:"attend"`
}

// PeopleResponse contains all users from database information
//
// swagger:response attendancesResponse
type attendancesResponse struct {
	// in: body
	Payload *[]Attend `json:"attend"`
}
