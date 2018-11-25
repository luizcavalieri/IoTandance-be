package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"

	"github.com/luizcavalieri/IoTandance-be/global"
)

var Db *sql.DB

func DbInit() {

	pgHost := os.Getenv("DB_HOST")
	pgPort := os.Getenv("DB_PORT")
	pgDbName := os.Getenv("DB_NAME")
	pgUser := os.Getenv("DB_USERNAME")
	pgPassword := os.Getenv("DB_PASSWORD")
	pgSslMode := os.Getenv("DB_SSL_MODE")
	connectUrl := "postgres://"+pgUser+":"+pgPassword+"@"+pgHost+":"+pgPort+"/"+pgDbName+"?sslmode="+pgSslMode
	log.Println("connStr: "+connectUrl)

	pgUrl, err := pq.ParseURL(connectUrl)

	//connStr := "host=%s port=%s dbname=%s user=%s sslmode=%s password=%s"
	//connStr = fmt.Sprintf(connStr, host, port, dbname, username, sslmode, password)
	//log.Println("connStr: "+connStr)


	//for _, pair := range os.Environ() {
	//	log.Println(pair)
	//}

	Db, err := sql.Open("postgres", pgUrl)
	log.Println("Checking if db is responding")
	global.LogFatal(err)

	err = Db.Ping()
	global.LogFatal(err)

	log.Println("Db successfully connected!")

}