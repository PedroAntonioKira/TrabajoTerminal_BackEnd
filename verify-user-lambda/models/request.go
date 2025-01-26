// models/request.go
package models

// ConfirmRequest estructura para almacenar los datos de la solicitud de confirmación.
type ConfirmRequest struct {
	Username         string `json:"username"`
	ConfirmationCode string `json:"confirmationCode"`
}
