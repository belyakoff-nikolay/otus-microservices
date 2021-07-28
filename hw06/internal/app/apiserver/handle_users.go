package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/belyakoff-nikolay/otus-microservices/hw06/internal/app/model"
)

func (s *APIServer) handleUserCreate() http.HandlerFunc {
	type UserCreateRequest struct {
		Email     string
		FirstName string
		LastName  string
	}

	type UserCreateResponse struct {
		ID    int64  `json:",omitempty"`
		Error string `json:"error,omitempty"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		var req UserCreateRequest
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			s.respondError(writer, err, http.StatusUnprocessableEntity)
			return
		}

		u := model.User{
			Email:     req.Email,
			FirstName: req.FirstName,
			LastName:  req.LastName,
		}
		newUser, err := s.storage.Users().Create(&u)
		if err != nil {
			s.respondError(writer, err, http.StatusInternalServerError)
			return
		}

		response := &UserCreateResponse{
			ID: newUser.ID,
		}
		s.respond(writer, response, http.StatusOK)
	}
}

func (s *APIServer) handleUserGet() http.HandlerFunc {
	type UserGetResponse struct {
		ID        int64  `json:",omitempty"`
		Email     string `json:",omitempty"`
		FirstName string `json:",omitempty"`
		LastName  string `json:",omitempty"`
		Error     string `json:"error,omitempty"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		userID, err := strconv.ParseInt(vars["userid"], 10, 64)
		if err != nil {
			s.logger.Debug(err)
			s.respondError(writer, err, http.StatusBadRequest)
			return
		}

		u, err := s.storage.Users().GetByID(userID)
		if err != nil {
			s.logger.Error(err)
			s.respondError(writer, err, http.StatusInternalServerError)
			return
		}

		response := &UserGetResponse{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}
		s.respond(writer, response, http.StatusOK)
	}
}

func (s *APIServer) handleUserDelete() http.HandlerFunc {
	type UserDropResponse struct {
		Error string `json:"error,omitempty"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		userID, err := strconv.ParseInt(vars["userid"], 10, 64)
		if err != nil {
			s.logger.Debug(err)
			s.respondError(writer, err, http.StatusBadRequest)
			return
		}

		err = s.storage.Users().Drop(userID)
		if err != nil {
			s.respondError(writer, err, http.StatusInternalServerError)
			return
		}

		response := &UserDropResponse{}
		s.respond(writer, response, http.StatusOK)
	}
}

func (s *APIServer) handleUserUpdate() http.HandlerFunc {
	type UserUpdateRequest struct {
		ID        int64
		Email     string
		FirstName string
		LastName  string
	}

	type UserUpdateResponse struct {
		Error string `json:"error,omitempty"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		var req UserUpdateRequest
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			s.respondError(writer, err, http.StatusUnprocessableEntity)
			return
		}

		u := model.User{
			ID:        req.ID,
			Email:     req.Email,
			FirstName: req.FirstName,
			LastName:  req.LastName,
		}

		err := s.storage.Users().Update(&u)
		if err != nil {
			s.respondError(writer, err, http.StatusInternalServerError)
			return
		}

		response := &UserUpdateResponse{}
		s.respond(writer, response, http.StatusOK)
	}
}
