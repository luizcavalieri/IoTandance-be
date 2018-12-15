package users

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/luizcavalieri/IoTandance-be/driver"
	"github.com/luizcavalieri/IoTandance-be/global"
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
	global.LogFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName,
			&user.LastName, &user.RoleId, &user.LastAccess,
			&user.Password, &user.RoleCd, &user.Active)
		global.LogFatal(err)

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

	rows, err := driver.Db.Query("SELECT * from users where user_id="+usrId)
	global.LogFatal(err)

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName,
			&user.LastName, &user.RoleId, &user.LastAccess,
			&user.Password, &user.RoleCd, &user.Active)

		if user.ID == usrId {
			w.WriteHeader(http.StatusOK)
			// add a arbitraty pause of 1 second
			time.Sleep(1000 * time.Millisecond)
			if err := json.NewEncoder(w).Encode(user); err != nil {
				global.LogFatal(err)
			}
			return
		}
		global.LogFatal(err)

	}
	json.NewEncoder(w).Encode(user)

}

// create a new item
func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("Create users", params["id"])
	var person User
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	users = append(users, person)
	json.NewEncoder(w).Encode(users)
}

// Delete an item
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("Delete users", params["id"])
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(users)
	}
}
