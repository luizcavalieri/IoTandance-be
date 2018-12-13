package people

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/luizcavalieri/IoTandance-be/driver"
	"github.com/luizcavalieri/IoTandance-be/global"
)

type Controller struct {
}

var people []User

func GetPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("Get people")
	driver.DbInit()
	var person User
	people = []User{}

	rows, err := driver.Db.Query("SELECT * from users")
	global.LogFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&person.ID, &person.Username, &person.FirstName,
			&person.LastName, &person.RoleId, &person.LastAccess,
			&person.Password, &person.RoleCd, &person.Active)
		global.LogFatal(err)

		people = append(people, person)
	}

	json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usrId := params["id"]
	log.Println("Get user", usrId)

	driver.DbInit()
	var person User
	people = []User{}

	rows, err := driver.Db.Query("SELECT * from users where user_id="+usrId)
	global.LogFatal(err)

	for rows.Next() {
		err := rows.Scan(&person.ID, &person.Username, &person.FirstName,
			&person.LastName, &person.RoleId, &person.LastAccess,
			&person.Password, &person.RoleCd, &person.Active)
		global.LogFatal(err)

		people = append(people, person)
	}
	json.NewEncoder(w).Encode(people)
}

// create a new item
func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("Create people", params["id"])
	var person User
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("Delete people", params["id"])
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
