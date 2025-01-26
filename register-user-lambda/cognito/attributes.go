package cognito

import (
	"aws-lambda-cognito/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

// BuildUserAttributes construye una lista de atributos personalizados para un usuario en AWS Cognito.
// Recibe un objeto `Request` y genera un slice de `types.AttributeType` con los valores correspondientes.
func BuildUserAttributes(req models.Request) []types.AttributeType {
	// Lista inicial de atributos obligatorios.
	attributes := []types.AttributeType{
		{Name: aws.String("email"), Value: aws.String(req.Email)},              // Atributo de correo electrónico.
		{Name: aws.String("phone_number"), Value: aws.String(req.PhoneNumber)}, // Atributo de número de teléfono.
		{Name: aws.String("name"), Value: aws.String(req.Name)},                // Atributo de nombre completo.
	}

	// Agregar atributo de rol si está presente.
	if req.Role != nil {
		attributes = append(attributes, types.AttributeType{
			Name: aws.String("custom:rol"), Value: aws.String(*req.Role),
		})
	}

	// Agregar atributo de estado si está presente.
	if req.Status != nil {
		attributes = append(attributes, types.AttributeType{
			Name: aws.String("custom:estatus"), Value: aws.String(*req.Status),
		})
	}

	// Agregar atributo de fecha de creación si está presente.
	if req.CreationDate != nil {
		attributes = append(attributes, types.AttributeType{
			Name: aws.String("custom:fecha_creacion"), Value: aws.String(*req.CreationDate),
		})
	}

	// Agregar atributo de fecha de modificación si está presente.
	if req.ModificationDate != nil {
		attributes = append(attributes, types.AttributeType{
			Name: aws.String("custom:fecha_modificacion"), Value: aws.String(*req.ModificationDate),
		})
	}

	// Retorna la lista completa de atributos.
	return attributes
}
