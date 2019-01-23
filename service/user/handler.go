package user

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

// PeopleResponse contains the user created in database
//
// swagger:response createUsersResponse
type createUsersResponse struct {
	// in: body
	ID int `json:"id"`
}
