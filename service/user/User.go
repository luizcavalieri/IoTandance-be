package user

type User struct {
	UserID       string // `json:"user_id,omitempty"`
	Username     string // `json:"username,omitempty"`
	UserFname    string // `json:"user_fname,omitempty"`
	UserLname    string // `json:"user_lname,omitempty"`
	RoleId       int    // `json:"role_id,omitempty"`
	LastAccessTm string // `json:"last_access_tm,omitempty"`
	Password     string // `json:"password,omitempty"`
	RoleCd       string // `json:"role_cd,omitempty"`
	ActiveYN     bool   // `json:"active_yn,omitempty"`
}
type Address struct {
	City  string        // `json:"city,omitempty"`
	State string        // `json:"state,omitempty"`
}
