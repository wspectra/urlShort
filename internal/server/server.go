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
	s.router.HandleFunc("/get", s.handleGet()).Methods("GET")
	s.router.HandleFunc("/post", s.handlePost()).Methods("POST")
}

func (s *ApiServer) settingStore() {
	switch s.conf.Store {
	case "inmemory":
		s.Store = store.NewInMemory()
	default:
		log.Fatal().Msg("[CONFIG]: wrong store flag")
	}

}

type RequestStruct struct {
	Url string
}

func (s *ApiServer) handlePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqStruct := RequestStruct{}
		if err := json.NewDecoder(r.Body).Decode(&reqStruct); err != nil {
			log.Fatal().Msg(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		reqStruct.Url, _ = s.Store.PostInfo(reqStruct.Url)
		if _, err := w.Write([]byte(reqStruct.Url)); err != nil { ///надо ли это?????
			log.Error().Msg(err.Error())
		}
	}
}

func (s *ApiServer) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqStruct := RequestStruct{}
		if err := json.NewDecoder(r.Body).Decode(&reqStruct); err != nil {
			log.Error().Msg("[REQUEST]: incorrect request body")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var err error
		if reqStruct.Url, err = s.Store.GetInfo(reqStruct.Url); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err := w.Write([]byte(err.Error())); err != nil { ///надо ли это?????
				log.Error().Msg("[REQUEST]: incorrect request body")
			}
			return
		}

		if _, err := w.Write([]byte(reqStruct.Url)); err != nil { ///надо ли это?????
			log.Error().Msg(err.Error())
		}
	}
}
