package models

// Request define la estructura de datos necesaria para interactuar con AWS Cognito.
// Contiene los campos obligatorios y opcionales para registrar o actualizar un usuario.
type Request struct {
	Username         string  `json:"username"`                    // Nombre de usuario (obligatorio).
	Password         string  `json:"password"`                    // Contraseña del usuario (obligatorio).
	Email            string  `json:"email"`                       // Correo electrónico del usuario (obligatorio).
	PhoneNumber      string  `json:"phone_number"`                // Número de teléfono en formato internacional (obligatorio).
	Name             string  `json:"name"`                        // Nombre completo del usuario (obligatorio).
	Role             *string `json:"role,omitempty"`              // Rol del usuario (opcional).
	Status           *string `json:"status,omitempty"`            // Estado del usuario (opcional).
	CreationDate     *string `json:"creation_date,omitempty"`     // Fecha de creación del usuario (opcional).
	ModificationDate *string `json:"modification_date,omitempty"` // Fecha de última modificación (opcional).
}
