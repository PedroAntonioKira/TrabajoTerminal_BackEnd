// utils/logger.go
package utils

import "log"

// LogError registra los errores o mensajes de información en la consola.
func LogError(message string, err ...error) {
	if len(err) > 0 {
		log.Printf("%s: %v", message, err[0])
	} else {
		log.Println(message)
	}
}
