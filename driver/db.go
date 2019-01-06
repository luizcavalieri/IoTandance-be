package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"

	"github.com/luizcavalieri/IoTendance-be/global"
)

var Db *sql.DB

func DbInit() {
	pgHost := os.Getenv("DB_HOST")
	pgPort := os.Getenv("DB_PORT")
	pgDbName := os.Getenv("DB_NAME")
	pgUser := os.Getenv("DB_USERNAME")
	pgPassword := os.Getenv("DB_PASSWORD")
	pgSslMode := os.Getenv("DB_SSL_MODE")

	pgUrl, err := pq.ParseURL(
		"postgres://" + pgUser +
			":" + pgPassword +
			"@" + pgHost +
			":" + pgPort +
			"/" + pgDbName +
			"?sslmode=" + pgSslMode)

	global.LogFatal(err, "")

	Db, err = sql.Open("postgres", pgUrl)
	global.LogFatal(err, "Connection db")

	err = Db.Ping()
	global.LogFatal(err, "")

	log.Println("Db successfully connected!")

}
