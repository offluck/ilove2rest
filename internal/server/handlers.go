package server

import (
	"net/http"

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
}

func (s *Server) getUserHandler(w http.ResponseWriter, _ *http.Request) {
}

func (s *Server) postUserHandler(w http.ResponseWriter, _ *http.Request) {
}

func (s *Server) putUserHandler(w http.ResponseWriter, _ *http.Request) {
}

func (s *Server) deleteUserHandler(w http.ResponseWriter, _ *http.Request) {
}
