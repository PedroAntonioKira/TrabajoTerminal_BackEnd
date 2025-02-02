package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/cognito"
	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/models"
	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/utils"
	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/validators"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// handler es la función principal que recibe solicitudes HTTP de API Gateway y procesa la confirmación de usuario.
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody models.ConfirmRequest

	// 🟢 Verificar y registrar el cuerpo de la solicitud
	log.Println("📥 Request Body recibido:", request.Body)

	// Decodificar el JSON recibido
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		utils.LogError("❌ Error al decodificar JSON", err)
		errorResponse := map[string]string{"error": "Invalid request. Ensure 'username' and 'confirmationCode' are provided"}
		body, _ := json.Marshal(errorResponse)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			Body: string(body),
		}, nil
	}

	// 🟢 Validar los campos obligatorios antes de enviar a Cognito
	if !validators.ValidateConfirmRequest(reqBody) {
		utils.LogError("❌ Campos requeridos faltantes en la solicitud")
		errorResponse := map[string]string{"error": "Missing required fields: 'username' and 'confirmationCode'"}
		body, _ := json.Marshal(errorResponse)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			Body: string(body),
		}, nil
	}

	// Llamar a la función de Cognito para confirmar el usuario
	response, err := cognito.ConfirmUser(reqBody)
	if err != nil {
		utils.LogError("❌ Error en confirmación de usuario", err)
		return response, nil
	}

	// 🟢 Respuesta exitosa
	log.Println("✅ Usuario confirmado correctamente")
	return response, nil
}

// Punto de entrada de la aplicación Lambda
func main() {
	lambda.Start(handler)
}
