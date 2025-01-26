package validators

import (
	"aws-lambda-cognito/models"
	"fmt"
)

// ValidateRequest realiza una validación completa de la estructura `Request`.
// Retorna un error si alguno de los campos obligatorios es inválido.
func ValidateRequest(req models.Request) error {
	// Valida el formato del número de teléfono.
	if !ValidatePhoneNumber(req.PhoneNumber) {
		return fmt.Errorf("el formato del número de teléfono es inválido")
	}

	// Valida que el correo no esté vacío.
	if req.Email == "" {
		return fmt.Errorf("el email es requerido")
	}

	// Si todas las validaciones pasan, retorna nil.
	return nil
}
