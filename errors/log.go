package errors

import "log"

/*
	Pretty prints an err to user
*/

func CheckErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
