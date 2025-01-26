package validators

import (
	"aws-lambda-cognito/models"
	"fmt"
	"log"
)

func ValidateGeneral(req models.Request) error {
	// Validar campos obligatorios
	if req.Username == "" {
		return fmt.Errorf("el campo 'username' es obligatorio")
	}
	if req.Password == "" {
		return fmt.Errorf("el campo 'password' es obligatorio")
	}
	if req.Email == "" {
		return fmt.Errorf("el campo 'email' es obligatorio")
	}
	if req.PhoneNumber == "" {
		return fmt.Errorf("el campo 'phone_number' es obligatorio")
	}
	if req.Name == "" {
		return fmt.Errorf("el campo 'name' es obligatorio")
	}

	// Opcional: Imprimir los campos opcionales si están presentes
	if req.Role != nil {
		log.Printf("Rol recibido: %s", *req.Role)
	}
	if req.Status != nil {
		log.Printf("Estado recibido: %s", *req.Status)
	}

	return nil
}
