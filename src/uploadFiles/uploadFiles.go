package main

import (
	"bytes"
	"encoding/json"
	"filesharing/mongoApi"
	"log"
	"mime"
	"mime/multipart"

	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) initSession() (err error) {
	s.session, err = mongoApi.Connect()
	return
}

func (s *service) uploadFiles(l logger.Logger, rbody *suckhttp.Request) (*suckhttp.Response, error) {
	_, params, _ := mime.ParseMediaType(rbody.GetHeader("content-type"))

	r := bytes.NewReader(rbody.Body)
	mr := multipart.NewReader(r, params["boundary"])
	form, _ := mr.ReadForm(100 << 20)

	files, ok := form.File["files"]
	if !ok {
		return suckhttp.NewResponse(400, "Bad Request"), nil
	}

	testDB := s.session.DB("filesharing")
	bucket := testDB.GridFS("tmpPrefix") // need to change

	var resp []struct {
		Filename string      `json:"filename"`
		Id       interface{} `json:"id"`
	}
	for _, fileHeader := range files {
		id, err := mongoApi.InsertFromMultipart(bucket, fileHeader)
		if err != nil {
			log.Println("Error uploadFile:", err)
			l.Error("[UploadFile]", err)
			return nil, err
		}
		resp = append(resp, struct {
			Filename string      "json:\"filename\""
			Id       interface{} "json:\"id\""
		}{Filename: fileHeader.Filename, Id: id})
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		l.Error("[Marshal response]", err)
		return nil, err
	}

	return suckhttp.NewResponse(200, "OK").AddHeader("Content-Type", "application/json").SetBody(respBytes), nil
}
