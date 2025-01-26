package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// ConfirmRequest estructura de la solicitud
type ConfirmRequest struct {
	Username         string `json:"username"`
	ConfirmationCode string `json:"confirmationCode"`
}

// Response estructura de la respuesta
type Response struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// Handler de la Lambda
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody ConfirmRequest

	// Deserializar JSON directamente
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil || reqBody.Username == "" || reqBody.ConfirmationCode == "" {
		log.Println("Error decoding JSON")
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       `{"error": "Invalid request. 'username' and 'confirmationCode' are required"}`,
		}, nil
	}

	// Crear sesión con AWS
	sess := session.Must(session.NewSession())
	cognito := cognitoidentityprovider.New(sess)

	// Confirmar registro en Cognito
	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String("6ilpp8jf9m70ged7fmu412nf3n"),
		Username:         aws.String(reqBody.Username),
		ConfirmationCode: aws.String(reqBody.ConfirmationCode),
	}

	_, err = cognito.ConfirmSignUp(input)
	if err != nil {
		log.Printf("Failed to confirm sign-up: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       fmt.Sprintf(`{"error": "Failed to confirm sign-up: %s"}`, err.Error()),
		}, nil
	}

	// Respuesta exitosa
	response := Response{
		Message: "User confirmed successfully",
	}

	responseBody, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}

func main() {
	lambda.Start(handler)
}
