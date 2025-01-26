// models/response.go
package models

// Response estructura para enviar respuestas JSON al cliente.
type Response struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
