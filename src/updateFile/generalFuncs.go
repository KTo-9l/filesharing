package main

import (
	"filesharing/mongoApi"
)

func (s *service) initSession(dbName, prefix string) (err error) {
	s.session, err = mongoApi.Connect()
	if err != nil {
		return err
	}
	s.bucket = s.session.DB(dbName).GridFS(prefix)
	return
}
