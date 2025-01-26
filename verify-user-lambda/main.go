// main.go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/cognito"
	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/models"
	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/validators"

	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// handler es la función principal que recibe solicitudes HTTP de API Gateway y procesa la confirmación de usuario.
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody models.ConfirmRequest

	// Decodifica el cuerpo JSON de la solicitud en la estructura reqBody
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil || !validators.ValidateConfirmRequest(reqBody) {
		utils.LogError("Invalid request")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "Invalid request. 'username' and 'confirmationCode' are required"}`,
		}, nil
	}

	// Llama a la función que confirma al usuario en Cognito
	err = cognito.ConfirmUser(reqBody)
	if err != nil {
		utils.LogError("Failed to confirm user", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "Failed to confirm user"}`,
		}, nil
	}

	// Respuesta exitosa en formato JSON
	response := models.Response{Message: "User confirmed successfully"}
	responseBody, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
	}, nil
}

// Punto de entrada principal de la aplicación Lambda
func main() {
	lambda.Start(handler)
}
