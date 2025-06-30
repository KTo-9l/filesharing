package main

import (
	"encoding/json"

	"filesharing/mongoApi"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) initSession() (err error) {
	s.session, err = mongoApi.Connect()
	return
}

func (s *service) getAllFiles(l logger.Logger) ([]byte, error) {
	testDB := s.session.DB("filesharing")
	bucket := testDB.GridFS("tmpPrefix") // need to change

	files, err := mongoApi.GetAllFiles(bucket)
	if err != nil {
		l.Error("[getAllFiles get query error]", err)
		return nil, err
	}

	respBytes, err := json.Marshal(files)
	if err != nil {
		l.Error("[getAllFiles error marshal files]", err)
		return nil, err
	}

	return respBytes, nil
}
