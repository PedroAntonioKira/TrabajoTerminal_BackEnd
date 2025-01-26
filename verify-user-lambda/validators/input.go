// validators/input.go
package validators

import "github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/models"

// ValidateConfirmRequest verifica si los campos obligatorios están presentes.
func ValidateConfirmRequest(req models.ConfirmRequest) bool {
	return req.Username != "" && req.ConfirmationCode != ""
}
