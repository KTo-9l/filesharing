package mongoApi

import (
	"log"

	"github.com/big-larry/mgo"
)

func Connect() (*mgo.Session, error) {
	uri := "mongodb://localhost:27017"
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}
	return session, nil
}
