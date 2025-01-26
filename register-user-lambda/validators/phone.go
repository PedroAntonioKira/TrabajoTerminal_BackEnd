package validators

import "regexp"

// ValidatePhoneNumber valida si un número de teléfono cumple con el formato internacional.
// El formato esperado comienza con "+" seguido de 1 a 15 dígitos.
func ValidatePhoneNumber(phoneNumber string) bool {
	re := regexp.MustCompile(`^\+[1-9]\d{1,14}$`) // Expresión regular para validar el número.
	return re.MatchString(phoneNumber)            // Retorna true si el formato es válido.
}
