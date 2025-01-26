package cognito

import (
	"aws-lambda-cognito/models"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

// RegisterUser registra un usuario en AWS Cognito utilizando el SDK.
// Recibe:
// - ctx: Contexto para la operación.
// - svc: Cliente de Cognito configurado.
// - req: Estructura `Request` con los datos del usuario.
// - userAttributes: Atributos personalizados del usuario.
// Retorna un error si falla la operación.
func RegisterUser(ctx context.Context, svc *cognitoidentityprovider.Client, req models.Request, userAttributes []types.AttributeType) error {
	// Configuración del input para la operación de registro.
	input := &cognitoidentityprovider.SignUpInput{
		ClientId:       aws.String("6ilpp8jf9m70ged7fmu412nf3n"), // ID del cliente de la app en Cognito.
		Username:       aws.String(req.Username),                 // Nombre de usuario.
		Password:       aws.String(req.Password),                 // Contraseña.
		UserAttributes: userAttributes,                           // Atributos personalizados del usuario.
	}

	// Llama al servicio de Cognito para registrar al usuario.
	_, err := svc.SignUp(ctx, input)
	if err != nil {
		// Retorna un error si la operación falla.
		return fmt.Errorf("error al registrar usuario en Cognito: %v", err)
	}

	// Retorna nil si la operación fue exitosa.
	return nil
}
