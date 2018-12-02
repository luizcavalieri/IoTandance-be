package user

type User struct {
	ID         string    `json:"user_id,omitempty"`
	Username   string    `json:"username,omitempty"`
	FirstName  string    `json:"user_fname,omitempty"`
	LastName   string    `json:"user_lname,omitempty"`
	RoleId     int       `json:"role_id,omitempty"`
	LastAccess string    `json:"last_access_tm,omitempty"`
	Password   string    `json:"password,omitempty"`
	RoleCd     string    `json:"role_cd,omitempty"`
	Active     bool      `json:"active_yn,omitempty"`
}
type Address struct {
	City  string        `json:"city,omitempty"`
	State string        `json:"state,omitempty"`
}
