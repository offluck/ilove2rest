package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/offluck/ilove2rest/internal/entities/user"
	"go.uber.org/zap"
)

func (s *Server) healthHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("I am healthy as an ox!"))
	if err != nil {
		s.logger.Error("Failed to write response", zap.Error(err))
	}
	s.logger.Debug("Sent heatlh check")
}

func (s *Server) getUsersHandler(w http.ResponseWriter, _ *http.Request) {
	users, err := s.DBClient.GetUsers(context.TODO())
	if err != nil {
		if err == sql.ErrNoRows {
			users = make([]user.UserDB, 0)
		} else {
			s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get users from DB: %+v", err))
		}
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to marshal users: %+v", err))
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonUsers)
	if err != nil {
		s.logger.Error("Failed to send users response", zap.Error(err))
	}
}

func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	user, err := s.DBClient.GetUser(context.TODO(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			s.writeError(w, http.StatusNotFound, fmt.Sprintf("Failed to find user %s", username))
		} else {
			s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get user %s from DB: %+v", username, err))
		}
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to marshal user %s: %+v", username, err))
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonUser)
	if err != nil {
		s.logger.Error("Failed to send user response", zap.String("username", username), zap.Error(err))
	}
}

func (s *Server) postUserHandler(w http.ResponseWriter, _ *http.Request) {
}

func (s *Server) putUserHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get("username")
}

func (s *Server) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get("username")
}

func (s *Server) writeError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(message))
	if err != nil {
		s.logger.Error("Failed to write error response", zap.Error(err))
	}
}
