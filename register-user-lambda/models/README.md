# Carpeta `models`

Esta carpeta contiene las estructuras de datos utilizadas en el proyecto. Estas estructuras definen los datos necesarios para las operaciones, validaciones e interacción con AWS Cognito.

## Archivos principales

- **`request.go`:** Define la estructura `Request`, que contiene los datos requeridos para manejar las solicitudes de registro y actualización de usuarios en Cognito.

## Estructuras disponibles

### **`Request`**
```go
type Request struct {
    Username         string  `json:"username"`          // Nombre de usuario (obligatorio)
    Password         string  `json:"password"`          // Contraseña del usuario (obligatorio)
    Email            string  `json:"email"`             // Correo electrónico (obligatorio)
    PhoneNumber      string  `json:"phone_number"`      // Número de teléfono (obligatorio, en formato internacional)
    Name             string  `json:"name"`              // Nombre completo del usuario (obligatorio)
    Role             *string `json:"role,omitempty"`    // Rol del usuario (opcional)
    Status           *string `json:"status,omitempty"`  // Estado del usuario (opcional)
    CreationDate     *string `json:"creation_date,omitempty"`  // Fecha de creación (opcional)
    ModificationDate *string `json:"modification_date,omitempty"`  // Fecha de modificación (opcional)
}
