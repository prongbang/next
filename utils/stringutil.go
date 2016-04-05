package utils

import (
	"log"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func String2Bytes(s string) []byte {
	byteArray := []byte(s)
	return byteArray
}
