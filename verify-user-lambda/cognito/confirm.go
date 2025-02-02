package cognito

import (
	"encoding/json"
	//"errors"
	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/models"
	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// clientID representa el identificador de cliente de Cognito, debe ser almacenado de forma segura.
const clientID = "6ilpp8jf9m70ged7fmu412nf3n"

// ConfirmUser procesa la confirmación del usuario en Cognito.
func ConfirmUser(req models.ConfirmRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSession())
	cognito := cognitoidentityprovider.New(sess)

	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(clientID),
		Username:         aws.String(req.Username),
		ConfirmationCode: aws.String(req.ConfirmationCode),
	}

	_, err := cognito.ConfirmSignUp(input)
	if err != nil {
		utils.LogError("❌ Error al confirmar usuario en Cognito", err)
		errorResponse := map[string]string{"error": "Failed to confirm user"}
		body, _ := json.Marshal(errorResponse)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Content-Type":                "application/json",
				"Access-Control-Allow-Origin": "*",
			},
			Body: string(body),
		}, nil
	}

	// 📌 Siempre devolver una respuesta JSON válida
	successResponse := map[string]string{"message": "User confirmed successfully"}
	body, _ := json.Marshal(successResponse)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(body),
	}, nil
}
