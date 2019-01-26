package user

/*
Params:

	      - name: user_fname
			type: string
			in: query
			maxItems: 10
			minItems: 3
			unique: false
	      - name: user_lname
			type: string
			in: query
			maxItems: 10
			minItems: 3
			unique: false
	      - name: role_id
			type: int
			in: query
			maxItems: 10
			minItems: 3
			unique: false
	      - name: password
			type: string
			in: query
			maxItems: 10
			minItems: 3
			unique: false
	      - name: role_cd
			type: string
			in: query
			maxItems: 10
			minItems: 3
			unique: false
	      - name: active_yn
			type: boolean
			in: query
			maxItems: 10
			minItems: 3
			unique: false

*/
// User description.
// swagger:model user
type User struct {
	// ID of the user
	//
	// required: true
	ID string `json:"user_id,omitempty"`
	// User name of the user
	//
	// required: true
	Username string `json:"username,omitempty"`
	// FirstName of the user
	//
	// required: false
	FirstName string `json:"user_fname,omitempty"`
	// LastName of the user
	//
	// required: false
	LastName string `json:"user_lname,omitempty"`
	// RoleId of the user
	//
	// required: false
	RoleId int `json:"role_id,omitempty"`
	// LastAccess of the user
	//
	// required: false
	LastAccess string `json:"last_access_tm,omitempty"`
	// Password of the user
	//
	// required: true
	Password string `json:"password,omitempty"`
	// RoleCd of the user
	//
	// required: false
	RoleCd string `json:"role_cd,omitempty"`
	// Active of the user
	//
	// required: true
	Active bool `json:"active_yn,omitempty"`
}

type LoginUser struct {
	// Username of the user
	//
	// required: true
	Username string `json:"username"`
	// Password of the user
	//
	// required: true
	Password string `json:"password"`
}

type LoginResponse struct {
}

func (logInUsr *LoginUser) getJwtToken() string {
	return logInUsr.Password + logInUsr.Username
}
