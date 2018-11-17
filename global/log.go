package global

import "log"

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

