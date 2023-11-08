package server

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type responseError struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}

func newResponseError(statusCode int, errorMessage string) responseError {
	return responseError{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}
}

func (s *Server) writeError(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)
	bytes, err := json.Marshal(newResponseError(statusCode, errorMessage))
	if err != nil {
		s.logger.Error("Failed to marshal response error", zap.Error(err))
	}

	_, err = w.Write(bytes)
	if err != nil {
		s.logger.Error("Failed to write error response", zap.Error(err))
	}
}
