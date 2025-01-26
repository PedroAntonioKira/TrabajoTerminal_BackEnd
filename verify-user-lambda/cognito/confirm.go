// cognito/confirm.go
package cognito

import (
	"errors"

	"github.com/PedroAntonioKira/TrabajoTerminal_BackEnd/verify-user-lambda/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// clientID representa el identificador de cliente de Cognito, debe ser almacenado de forma segura.
const clientID = "6ilpp8jf9m70ged7fmu412nf3n"

// ConfirmUser procesa la confirmación del usuario en Cognito.
func ConfirmUser(req models.ConfirmRequest) error {
	sess := session.Must(session.NewSession())
	cognito := cognitoidentityprovider.New(sess)

	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(clientID),
		Username:         aws.String(req.Username),
		ConfirmationCode: aws.String(req.ConfirmationCode),
	}

	_, err := cognito.ConfirmSignUp(input)
	if err != nil {
		return errors.New("failed to confirm user")
	}
	return nil
}
