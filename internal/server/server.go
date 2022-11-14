package server

import "github.com/wspectra/urlShort/internal/config"

type ApiServer struct {
	conf *config.Config
}

func NewServer() *ApiServer {
	return &ApiServer{
		conf: config.NewConfig(),
	}
}

func (s *ApiServer) Start() error {
	return nil
}
