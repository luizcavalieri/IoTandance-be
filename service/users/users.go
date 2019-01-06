package users

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
// swagger:parameters listUser
type IDParam struct {
	// The ID of a user
	//
	// in: path
	// required: true
	ID string `json:"user_id"`
}

var users []User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /people people listPeople
	//
	// Lists all users.
	//
	// This will show all recorded people.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200: usersResponse

	log.Println("Get users")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	driver.DbInit()
	var user User
	users = []User{}

	rows, err := driver.Db.Query("SELECT * from users")
	global.LogFatal(err, "")

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName,
			&user.LastName, &user.RoleId, &user.LastAccess,
			&user.Password, &user.RoleCd, &user.Active)
		global.LogFatal(err, "No users found")

		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// Display a single data
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)
	usrId := params["id"]
	log.Println("Get user", usrId)

	driver.DbInit()
	var user User

	rows, err := driver.Db.Query("SELECT * from users where user_id=" + usrId)
	global.LogFatal(err, "No user found with the id")

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName,
			&user.LastName, &user.RoleId, &user.LastAccess,
			&user.Password, &user.RoleCd, &user.Active)

		if user.ID == usrId {
			w.WriteHeader(http.StatusOK)
			// add a arbitraty pause of 1 second
			time.Sleep(1000 * time.Millisecond)
			if err := json.NewEncoder(w).Encode(user); err != nil {
				global.LogFatal(err, "")
			}
			return
		}
		global.LogFatal(err, "")

	}
	json.NewEncoder(w).Encode(user)

}

// Display a single data
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /users users createUsers
	//
	// Create user from username, user_fname, user_lname, role_id, password, role_cd, active_yn.
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

	log.Println("Create user")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	driver.DbInit()
	var user User
	var userId int
	json.NewDecoder(r.Body).Decode(&user)
	log.Println("Create user", user.Username)

	// Verifies if username is already taken

	var count int

	driver.Db.QueryRow(
		"SELECT COUNT(user_id) FROM users where username=$1;", user.Username).Scan(&count)

	if count > 0 {
		log.Println("Username already exists: ", user.Username)
		w.WriteHeader(http.StatusConflict)

	} else {
		roleId := 1
		roleCd := "test"
		activeYn := true

		err := driver.Db.QueryRow(
			"Insert into "+
				"users (username, user_fname, user_lname, role_id, password, role_cd, active_yn)"+
				"values($1, $2, $3, $4, $5, $6, $7) Returning user_id; ",
			user.Username,
			user.FirstName,
			user.LastName,
			roleId,
			user.Password,
			roleCd,
			activeYn).Scan(&userId)

		global.LogFatal(err, "Error after insert into create user")

		json.NewEncoder(w).Encode(userId)
		w.WriteHeader(http.StatusOK)
	}

	return
}
