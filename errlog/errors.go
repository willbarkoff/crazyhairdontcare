package errlog

import "log"

//LogError logs the error
func LogError(context string, err error) {
	log.Println("An error occured " + context)
	log.Println(err)
}
