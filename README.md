Confirmación de Usuario en AWS Cognito

Este proyecto es una AWS Lambda en Go que permite la confirmación de usuarios en AWS Cognito mediante API Gateway.

Estructura del Proyecto

├── cognito
│   └── confirm.go   # Lógica para interactuar con Cognito
├── models
│   ├── request.go   # Definición de la estructura de la solicitud
│   └── response.go  # Definición de la estructura de la respuesta
├── utils
│   └── logger.go    # Función para registro de errores
├── validators
│   └── input.go     # Validaciones de entrada
├── main.go          # Punto de entrada principal
├── go.mod           # Gestión de módulos de Go
├── go.sum           # Dependencias del proyecto
├── package_lambda.bat  # Script para empaquetado de la Lambda
└── README.md        # Documentación

Requisitos Previos

Tener una cuenta de AWS con permisos para Lambda y API Gateway.

Tener instalado Go (versión 1.18 o superior).

Configurar AWS CLI con las credenciales adecuadas.

Instalación

Clonar el repositorio:

git clone https://github.com/tu-repositorio/confirm-user-lambda.git
cd confirm-user-lambda

Instalar dependencias:

go mod tidy

Configuración

Configurar la variable de entorno para el Client ID de Cognito:

export COGNITO_CLIENT_ID=tu-client-id

Modificar el archivo cognito/confirm.go para obtener el Client ID desde la variable de entorno:

clientID := os.Getenv("COGNITO_CLIENT_ID")

Despliegue

Para empaquetar y subir la Lambda a AWS, ejecutar:

sh package_lambda.bat

Luego, subir el archivo main.zip a AWS Lambda.

Pruebas

Crear un endpoint en API Gateway apuntando a la Lambda.

Probar la API utilizando Postman con la siguiente configuración:

URL: https://your-api-id.execute-api.us-east-1.amazonaws.com/dev/confirm-user

Método: POST

Encabezados:

Content-Type: application/json

Cuerpo:

{
  "username": "prueba01@mailinator.com",
  "confirmationCode": "244678"
}

Funcionamiento

main.go inicia la ejecución de la Lambda y delega la solicitud al paquete cognito.

cognito/confirm.go maneja la lógica de confirmación de usuario en AWS Cognito.

validators/input.go verifica que los datos de entrada sean correctos.

utils/logger.go maneja el registro de errores.

Consideraciones de Seguridad

Se recomienda almacenar el Client ID de Cognito como una variable de entorno en AWS Lambda.

Asegurarse de que los permisos de la Lambda permitan interactuar con Cognito.

Autor

Tu Nombre - GitHub

Licencia

Este proyecto está bajo la licencia MIT - consulta el archivo LICENSE para más detalles.

