package handler

import (
	"encoding/json"
	"net/http"
)

// ResponseFactory - фабрика для создания ответов
type ResponseFactory struct{}

// NewResponseFactory - создает новый экземпляр ResponseFactory
func NewResponseFactory() *ResponseFactory {
	return &ResponseFactory{}
}

// SendJSON - отправляет JSON-ответ
func (rf *ResponseFactory) SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

type Error struct {
	Error string `json:"error"`
}

// SendError - отправляет JSON-ответ с ошибкой
func (rf *ResponseFactory) SendError(w http.ResponseWriter, statusCode int, message string) {
	rf.SendJSON(w, statusCode, Error{Error: message})
}

// SendSuccess - отправляет успешный JSON-ответ
func (rf *ResponseFactory) SendSuccess(w http.ResponseWriter, statusCode int, message string) {
	rf.SendJSON(w, statusCode, map[string]string{"message": message})
}
