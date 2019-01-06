package global

import "log"

func LogFatal(err error, msg string) {
	if err != nil {
		if &msg != nil {
			log.Println(msg)
		}
		log.Fatal(err)
	}
}
