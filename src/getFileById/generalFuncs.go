package main

import (
	"filesharing/mongoApi"
	"io"

	"github.com/big-larry/mgo"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) initSession(dbName, prefix string) (err error) {
	s.session, err = mongoApi.Connect()
	if err != nil {
		return err
	}
	s.bucket = s.session.DB(dbName).GridFS(prefix)
	return
}

func fileToBytes(l logger.Logger, file *mgo.GridFile) ([]byte, error) {
	bytes, err := io.ReadAll(file)
	if err != nil {
		l.Error("[error reading file]", err)
		return nil, err
	}
	return bytes, nil
}
