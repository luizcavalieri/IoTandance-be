package global

import (
	"database/sql"
	"fmt"
	"log"
)

func LogFatal(err error, txt string) {
	if err != nil {
		log.Println(txt)
		log.Fatal(err)
	}
}

func LogAndQuery(db *sql.DB, query string, args ...interface{}) *sql.Rows {
	fmt.Println(query)
	res, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	return res
}

