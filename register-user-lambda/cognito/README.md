# Carpeta `cognito`

Esta carpeta contiene las funciones relacionadas con la interacción con el servicio AWS Cognito. Aquí se manejan las operaciones necesarias para registrar usuarios, construir atributos y realizar acciones específicas con el servicio.

## Archivos principales

- **`attributes.go`:** Define y construye los atributos personalizados necesarios para los usuarios en Cognito.
- **`register.go`:** Registra un usuario en Cognito utilizando las credenciales proporcionadas.

## Uso

1. **Construir atributos:**
   Usa `BuildUserAttributes` para construir los atributos personalizados antes de registrar un usuario.
   ```go
   attributes := BuildUserAttributes(request)
