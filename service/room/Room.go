package room

// Room description.
// swagger:model room
type Room struct {
	// ID of the room
	//
	// required: true
	ID int `json:"room_id,omitempty"`
	// Name room
	//
	// required: true
	Name string `json:"name,omitempty"`
	// Room description
	//
	// required: false
	Description string `json:"description,omitempty"`
	// Room capacity
	//
	// required: true
	Capacity int `json:"capacity,omitempty"`
	// Sort
	//
	// required: true
	Sort int `json:"sort,omitempty"`
	// Removed
	//
	// required: true
	Removed bool `json:"removed,omitempty"`
}
