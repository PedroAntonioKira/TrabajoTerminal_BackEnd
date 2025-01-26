package utils

import (
	"log"
	"reflect"
)

// DebugStruct imprime los detalles de una estructura, incluyendo sus campos y valores.
// Útil para depurar estructuras complejas durante el desarrollo.
func DebugStruct(v interface{}) {
	val := reflect.ValueOf(v) // Obtiene el valor de la estructura.
	typ := reflect.TypeOf(v)  // Obtiene el tipo de la estructura.

	log.Println("---- Detalle de campos y valores recibidos ----")
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i) // Campo actual.
		value := val.Field(i) // Valor del campo.

		// Verifica si el campo es un puntero y está vacío.
		if value.Kind() == reflect.Ptr && value.IsNil() {
			log.Printf("Campo: %s (JSON: %s), Valor: <nil>", field.Name, field.Tag.Get("json"))
		} else {
			log.Printf("Campo: %s (JSON: %s), Valor: %v", field.Name, field.Tag.Get("json"), value.Interface())
		}
	}
	log.Println("---------------------------------------------")
}
