package user

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/luizcavalieri/IoTendance-be/driver"
	"github.com/luizcavalieri/IoTendance-be/global"
)

// IDParam is used to identify a user
//
// swagger:parameters users
type IDParam struct {
	// The ID of a user
	//
	// in: path
	// required: true
	ID string `json:"user_id"`
}

var users []User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	/* swagger:route GET /users people listPeople

	   Lists all users.

	   This will show all recorded people.

	     Consumes:
	     - application/json

	     Produces:
	     - application/json

	     Schemes: http, https

	     Responses:
		   200: usersResponse
	*/

	log.Println(`Get users`)

	w.Header().Set(`Content-Type`, `application/json; charset=UTF-8`)
	w.WriteHeader(http.StatusOK)

	driver.DbInit()
	var user User
	users = []User{}

	rows, err := driver.Db.Query(`SELECT * from users`)
	global.LogFatal(err, ``)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName,
			&user.LastName, &user.RoleId, &user.LastAccess,
			&user.Password, &user.RoleCd, &user.Active)
		global.LogFatal(err, `No users found`)

		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /users/{id} users listUsers
	//
	// Lists user from their id.
	//
	// This will show the record of an identified user.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Params:
	//     - id: IDParam
	//
	//     Responses:
	//       200: userResponse
	//       404: jsonError

	w.Header().Set(`Content-Type`, `application/json; charset=UTF-8`)

	params := mux.Vars(r)
	usrId := params[`id`]
	log.Println(`Get user`, usrId)

	driver.DbInit()
	var user User

	rows, err := driver.Db.Query(`SELECT * from users where user_id=$1`, usrId)
	global.LogFatal(err, `No user found with the id`)

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName,
			&user.LastName, &user.RoleId, &user.LastAccess,
			&user.Password, &user.RoleCd, &user.Active)

		if user.ID == usrId {
			w.WriteHeader(http.StatusOK)
			// add a arbitraty pause of 1 second
			time.Sleep(1000 * time.Millisecond)
			if err := json.NewEncoder(w).Encode(user); err != nil {
				global.LogFatal(err, ``)
			}
			return
		}
		global.LogFatal(err, ``)

	}
	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	/*

	 swagger:route POST /users users createUsers

	 Create new user.

	 This will create a record of an user.

	     Consumes:
	     - application/json

	     Produces:
	     - application/json

	     Schemes: http, https

	     Responses:
	       200: createUsersResponse
	       404: jsonError
	*/

	log.Println(`Create user`)
	w.Header().Set(`Content-Type`, `application/json; charset=UTF-8`)

	driver.DbInit()
	var user User
	var userId int
	json.NewDecoder(r.Body).Decode(&user)
	log.Println(`Create user`, user.Username)

	// Verifies if username is already taken

	var count int

	driver.Db.QueryRow(
		`SELECT COUNT(user_id) FROM users where username=$1;`, user.Username).Scan(&count)

	if count > 0 {
		log.Println(`Username already exists: `, user.Username)
		w.WriteHeader(http.StatusConflict)

	} else {
		roleId := 1
		roleCd := `test`
		activeYn := true

		err := driver.Db.QueryRow(
			`Insert into 
					users (username, user_fname, user_lname, role_id, password, role_cd, active_yn)
					values($1, $2, $3, $4, $5, $6, $7) Returning user_id; `,
			user.Username,
			user.FirstName,
			user.LastName,
			roleId,
			user.Password,
			roleCd,
			activeYn).Scan(&userId)

		global.LogFatal(err, `Error after insert into create user`)

		json.NewEncoder(w).Encode(userId)
		w.WriteHeader(http.StatusOK)
	}

	return
}

func Login(w http.ResponseWriter, r *http.Request) {

	/*

	 swagger:route POST /login users Login

	 This will verify if there is user with the combination of password and username.

	 This will login the user into the system.

	     Consumes:
	     - application/json

	     Produces:
	     - application/json

	     Schemes: http, https

	     Responses:
	       200: getUserResponse
	       404: jsonError
	*/

	w.Header().Set(`Content-Type`, `application/json; charset=UTF-8`)

	var usrLgn LoginUser
	json.NewDecoder(r.Body).Decode(&usrLgn)
	username := usrLgn.Username
	password := usrLgn.Password
	log.Println(`Get user `, usrLgn.Username)

	driver.DbInit()
	var user User

	rows, err := driver.Db.Query(
		`SELECT * from users where username='%s' and password='%s'`,
		username,
		password)
	global.LogFatal(err, `No user found with the combination of username and password`)

	defer driver.Db.Close()

	// Maps user
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName,
			&user.LastName, &user.RoleId, &user.LastAccess,
			&user.Password, &user.RoleCd, &user.Active)
		global.LogFatal(err, `Not able to map user`)

		// Confirms is the correct user
		if user.Username == username && user.Password == password {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(user)
			global.LogFatal(err, `Error on encode`)
		}
	}
}
