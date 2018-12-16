package attend

// JsonError is a generic error in JSON format
//
// swagger:response jsonError
type jsonError struct {
	// in: body
	Message string `json:"message"`
}

// UserResponse contains a single person information
//
// swagger:response userResponse
type userResponse struct {
	// in: body
	Payload *User `json:"user"`
}

// PeopleResponse contains all users from database information
//
// swagger:response usersResponse
type usersResponse struct {
	// in: body
	Payload *[]User `json:"users"`
}
