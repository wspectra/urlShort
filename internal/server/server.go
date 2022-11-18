package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/wspectra/urlShort/internal/config"
	"github.com/wspectra/urlShort/internal/store"
	"net/http"
)

type ApiServer struct {
	conf   *config.Config
	router *mux.Router
	Store  store.Store
}

func NewServer() *ApiServer {
	return &ApiServer{
		conf:   config.NewConfig(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {
	s.configureRouter()
	log.Info().Msg("starting api server on  " + s.conf.BindPort)
	s.settingStore()
	return http.ListenAndServe(s.conf.BindPort, s.router)
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *ApiServer) settingStore() {
	switch s.conf.Store {
	case "inmemory":
		s.Store = store.NewInMemory()
	default:
		log.Fatal().Msg("[CONFIG]: wrong store flag")
	}

}

type Url struct {
	Url string
}

func (s *ApiServer) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		url := Url{}
		json.NewDecoder(r.Body).Decode(&url)
		switch r.Method {
		case "POST":
			url.Url = s.Store.PostInfo(url.Url)
			w.Write([]byte(url.Url))
		case "GET":
			url.Url = s.Store.GetInfo(url.Url)
			w.Write([]byte(url.Url))
		default:
			log.Error().Msg("[REQUEST]: wrong request")
		}
	}
}
