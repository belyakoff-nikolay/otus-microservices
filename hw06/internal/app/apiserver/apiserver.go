package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/belyakoff-nikolay/otus-microservices/hw06/internal/pgstorage"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *pgstorage.Storage
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logConfig()

	s.logger.Infof("staring...")

	s.configureRouter()

	if err := s.configureStorage(); err != nil {
		return err
	}

	err := http.ListenAndServe(s.config.BindAddr, s.router)
	return err
}

func (s *APIServer) logConfig() {
	b, err := json.Marshal(s.config)
	if err != nil {
		s.logger.Error("can't marshal config:%v", err)
	} else {
		s.logger.Debugf("config: %v", string(b))
	}
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLever)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.Handle("/liveness", s.handleLiveness())
	s.router.Handle("/user", s.handleUserCreate()).Methods("POST")
	s.router.Handle("/user", s.handleUserUpdate()).Methods("PUT")
	s.router.Handle("/user/{userid}", s.handleUserGet()).Methods("GET")
	s.router.Handle("/user/{userid}", s.handleUserDelete()).Methods("DELETE")
}

func (s *APIServer) configureStorage() error {
	config := &pgstorage.Config{DatabaseURL: s.config.DatabaseURL}
	storage := pgstorage.New(config)
	err := storage.Open()
	if err != nil {
		return err
	}
	s.storage = storage
	return nil
}

func (s *APIServer) handleLiveness() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(`{"status": "OK"}`))
	}
}

func (s *APIServer) respond(writer http.ResponseWriter, data interface{}, code int) {
	b, err := json.Marshal(data)
	if err != nil {
		err = fmt.Errorf("marshal to json %T: %w", data, err)
		s.logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	_, err = writer.Write(b)
	if err != nil {
		s.logger.Debugf("can't write response:%v", err)
	}
}

func (s *APIServer) respondError(writer http.ResponseWriter, err error, code int) {
	s.respond(writer, map[string]string{"error": err.Error()}, code)
}
