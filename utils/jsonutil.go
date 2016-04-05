package utils

import (
	"encoding/json"
	"log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Type2JsonString(inf interface{}) string {
	b, err := json.Marshal(inf)
	if err != nil {
		log.Fatal(b)
	}
	return string(b)
}

func Type2JsonByte(inf interface{}) []byte {
	b, err := json.Marshal(inf)
	if err != nil {
		log.Fatal(b)
	}
	return b
}

func BindJSON(c *gin.Context, t interface{}) (error) {

	// read body
	b, _ := ioutil.ReadAll(c.Request.Body)

	// convert json byte to type
	err := json.Unmarshal(b, &t)

	return err
}
