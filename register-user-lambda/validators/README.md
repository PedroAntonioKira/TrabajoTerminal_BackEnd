
---

### **validators/README.md**
```markdown
# Carpeta `validators`

Esta carpeta contiene las funciones para validar datos de entrada antes de enviarlos a AWS Cognito. Garantiza que los datos cumplan con los requisitos necesarios.

## Archivos principales

- **`phone.go`:** Valida el formato del número de teléfono.
- **`request.go`:** Valida que todos los campos requeridos en el `Request` sean correctos.

## Uso

1. **Validar número de teléfono:**
   ```go
   isValid := ValidatePhoneNumber("+123456789")
   if !isValid {
       log.Println("Número de teléfono inválido")
   }
