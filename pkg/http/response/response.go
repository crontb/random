package response

import (
	"encoding/json"
	"net/http"
)

type JSON struct {
	Result interface{} `json:"result,omitempty"`
	Error  *JSONError  `json:"error,omitempty"`
}

type JSONError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (s *JSON) SetError(code int, message string) {
	s.Error = &JSONError{
		Code:    code,
		Message: message,
	}
}

func (s *JSON) SetResult(data interface{}) {
	s.Result = data
}

func (s *JSON) Write(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}
