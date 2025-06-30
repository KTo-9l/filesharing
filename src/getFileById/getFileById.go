package main

import (
	"fmt"
	"io"

	"filesharing/mongoApi"

	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) initSession() (err error) {
	s.session, err = mongoApi.Connect()
	return
}

func (s *service) getFileById(l logger.Logger, id string) (*suckhttp.Response, error) {
	testDB := s.session.DB("filesharing")
	bucket := testDB.GridFS("tmpPrefix") // need to change

	file, err := mongoApi.GetFileById(bucket, id)
	if err != nil {
		l.Error("[getFileById get query error]", err)
		return nil, err
	}

	respBytes, err := io.ReadAll(file)
	if err != nil {
		l.Error("[getAllFiles error reading file]", err)
		return nil, err
	}

	response := suckhttp.NewResponse(200, "OK").
		AddHeader("Content-Type", "application/octet-stream").
		AddHeader("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, file.Name())).
		SetBody(respBytes)

	return response, nil
}
