package main

import (
	"context"

	"github.com/big-larry/mgo"
	"github.com/big-larry/suckhttp"
	"github.com/okonma-violet/services/logs/logger"
	"github.com/okonma-violet/services/universalservice_nonepoll"
)

type config struct {
}

type service struct {
	session *mgo.Session
}

const thisServiceName universalservice_nonepoll.ServiceName = "filesharing_get_file_by_id"

func (c *config) InitFlags() {}

func (c *config) PrepareHandling(ctx context.Context, pubs_getter universalservice_nonepoll.Publishers_getter) (universalservice_nonepoll.BaseHandleFunc, universalservice_nonepoll.Closer, error) {
	s := &service{}

	var err error

	if err = s.initSession(); err != nil {
		return nil, nil, err
	}

	return universalservice_nonepoll.CreateHTTPHandleFunc(s), s, nil
}

func (s *service) HandleHTTP(req *suckhttp.Request, l logger.Logger) (response *suckhttp.Response, err error) {
	if req.GetMethod() == suckhttp.GET {
		id := req.Uri.Query().Get("id")
		if id == "" {
			response = suckhttp.NewResponse(400, "Bad Request")
			return
		}

		if resp, err := s.getFileById(l, id); err != nil {
			response = suckhttp.NewResponse(500, "Internal Server Error")
		} else {
			return resp, err
		}
	} else {
		response = suckhttp.NewResponse(405, "Method Not Allowed")
	}
	return
}

func (s *service) Close(l logger.Logger) error {
	s.session.Close()
	return nil
}

func main() {
	universalservice_nonepoll.InitNewService(thisServiceName, &config{}, 1)
}
