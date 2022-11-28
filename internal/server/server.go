package server

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/wspectra/urlShort/internal/config"
	"github.com/wspectra/urlShort/internal/store"
	"net/http"
	"net/url"
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
	s.configureStore()
	return http.ListenAndServe(s.conf.BindPort, s.router)
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/{shortUrl}", s.handleGet()).Methods("GET")
	s.router.HandleFunc("/post", s.handlePost()).Methods("POST")
}

func (s *ApiServer) configureStore() {
	log.Info().Msg("configuring store...")
	switch s.conf.Store {
	case "inmemory":
		s.Store = store.NewInMemory()
	case "postgres":
		st := store.NewPstStore(s.conf)
		if err := st.Open(); err != nil {
			log.Fatal().Msg(err.Error())
		}
		s.Store = st
		log.Info().Msg("[API-SERVER]: Successfuly connected to database")
	default:
		log.Fatal().Msg("[CONFIG]: wrong store flag")
	}

}

type RequestStruct struct {
	Url string `json:"Url" validate:"required"`
}

func (s *ApiServer) handlePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqStruct := RequestStruct{}

		//проверка JSON на ошибку декода
		if err := json.NewDecoder(r.Body).Decode(&reqStruct); err != nil {
			log.Error().Msg(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//проверка JSON на валидность JSON
		validate := validator.New()
		if err := validate.Struct(reqStruct); err != nil {
			log.Error().Msg(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//проверка ссылки на валидность
		if _, err := url.ParseRequestURI(reqStruct.Url); err != nil {
			log.Error().Msg(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var err error
		reqStruct.Url, err = s.Store.PostInfo(reqStruct.Url)
		if err != nil {
			log.Error().Msg(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write([]byte(reqStruct.Url)); err != nil {
			log.Error().Msg(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s *ApiServer) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		shortUrl := vars["shortUrl"]

		longUrl, err := s.Store.GetInfo(shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Error().Msg(err.Error())
			return
		}
		http.Redirect(w, r, longUrl, http.StatusPermanentRedirect)

	}
}
