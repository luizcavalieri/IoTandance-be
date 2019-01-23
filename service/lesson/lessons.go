package lesson

// IDParam is used to identify a user
//
// swagger:parameters users
type IDParam struct {
	// The ID of a user
	//
	// in: path
	// required: true
	ID string `json:"lesson_id"`
}
