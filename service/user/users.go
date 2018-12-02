package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/luizcavalieri/IoTandance-be/driver"
	"github.com/luizcavalieri/IoTandance-be/global"

)

type Controller struct {
}

var users []User

type userResolver struct {
	u *User
}

//func (_ *glob.Query) Hello() string { return "Hello, world!" }

func (r *helloWorldResolver) Hello(ctx context.Context) (string, error) {
	return "Hello world!", nil
}

func (_ *global.Query) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Get users")
	driver.DbInit()
	var person User
	users = []User{}

	rows, err := driver.Db.Query("SELECT * from users")
	global.LogFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&person.ID, &person.Username, &person.FirstName,
			&person.LastName, &person.RoleId, &person.LastAccess,
			&person.Password, &person.RoleCd, &person.Active)
		global.LogFatal(err)

		users = append(users, person)
	}

	json.NewEncoder(w).Encode(users)
}

// Display a single data
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("Get user", params["id"])
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
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
