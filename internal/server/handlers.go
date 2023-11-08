package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/offluck/ilove2rest/internal/entities/user"
	"go.uber.org/zap"
)

func (s *Server) healthHandler(w http.ResponseWriter, _ *http.Request) {
	s.logger.Debug("Health handler")

	_, err := w.Write([]byte("I am healthy as an ox!"))
	if err != nil {
		s.logger.Error("Failed to write response", zap.Error(err))
	}

	s.logger.Debug("Sent heatlh check")
}

func (s *Server) getUsersHandler(w http.ResponseWriter, _ *http.Request) {
	s.logger.Debug("Get users handler")

	usersDB, err := s.DBClient.GetUsers(context.TODO())
	if err != nil {
		if err == sql.ErrNoRows {
			usersDB = make([]user.UserDB, 0)
		} else {
			s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get users from DB: %+v", err))
			return
		}
	}

	usersResponse := make([]user.UserResponse, 0, len(usersDB))
	for _, userDB := range usersDB {
		usersResponse = append(usersResponse, userDB.DB2Resp())
	}

	jsonUsers, err := json.Marshal(usersResponse)
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to marshal users: %+v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonUsers)
	if err != nil {
		s.logger.Error("Failed to send users response", zap.Error(err))
	}

	s.logger.Debug("Sent users")
}

func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("Get user handler")

	username := chi.URLParam(r, "username")
	s.logger.Debug("Got username", zap.String("username", username))

	userDB, err := s.DBClient.GetUser(context.TODO(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			s.writeError(w, http.StatusNotFound, fmt.Sprintf("Failed to find user %s", username))
			return
		}

		s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get user %s from DB: %+v", username, err))
		return
	}

	jsonUser, err := json.Marshal(userDB.DB2Resp())
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to marshal user %s: %+v", username, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonUser)
	if err != nil {
		s.logger.Error("Failed to send user response", zap.String("username", username), zap.Error(err))
	}

	s.logger.Debug("Sent user")
}

func (s *Server) postUserHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("Post user handler")

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		s.writeError(w, http.StatusBadRequest, "Failed to read request")
		return
	}

	userRequest := user.UserRequest{}
	err = json.Unmarshal(bytes, &userRequest)
	if err != nil || !userRequest.IsValid() {
		s.writeError(w, http.StatusBadRequest, "Failed to parse JSON request")
		return
	}

	userDB, err := s.DBClient.AddUser(context.TODO(), userRequest.Req2DB())
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to add user to DB: %+v", err))
		return
	}

	jsonUser, err := json.Marshal(userDB.DB2Resp())
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to marshal user %s: %+v", *userRequest.Username, err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonUser)
	if err != nil {
		s.logger.Error("Failed to send user response", zap.String("username", *userRequest.Username), zap.Error(err))
	}

	s.logger.Debug("Created user")
}

func (s *Server) putUserHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("Put user handler")

	username := chi.URLParam(r, "username")
	s.logger.Debug("Got username", zap.String("username", username))

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		s.writeError(w, http.StatusBadRequest, "Failed to read request")
		return
	}

	userRequest := user.UserRequest{}
	err = json.Unmarshal(bytes, &userRequest)
	if err != nil || !userRequest.IsValid() {
		s.writeError(w, http.StatusBadRequest, "Failed to parse JSON request")
		return
	}

	userDB, err := s.DBClient.UpdateUser(context.TODO(), username, userRequest.Req2DB())
	if err != nil {
		if err == sql.ErrNoRows {
			s.writeError(w, http.StatusNotFound, fmt.Sprintf("Failed to find user %s", username))
			return
		}

		s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to put user %s from DB: %+v", username, err))
		return
	}

	jsonUser, err := json.Marshal(userDB.DB2Resp())
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to marshal user %s: %+v", *userRequest.Username, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonUser)
	if err != nil {
		s.logger.Error("Failed to send user response", zap.String("username", *userRequest.Username), zap.Error(err))
	}

	s.logger.Debug("Updated user")
}

func (s *Server) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("Delete user handler")

	username := chi.URLParam(r, "username")
	s.logger.Debug("Got username", zap.String("username", username))

	err := s.DBClient.DeleteUser(context.TODO(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			s.writeError(w, http.StatusNotFound, fmt.Sprintf("Failed to find user %s", username))
			return
		}
		s.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete user %s from DB: %+v", username, err))
		return
	}

	w.WriteHeader(http.StatusNoContent)

	s.logger.Debug("Deleted user")
}
