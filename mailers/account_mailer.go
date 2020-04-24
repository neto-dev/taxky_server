package mailers

import (
	"log"
	"os"

	"github.com/taxky/models"
)

// HOW TO

// Crear una carpeta y sus archivos html con el contenido del correo
// en la carpeta views/mailers

// el template usa la sintaxis {{ .Variable }} para sustituir la
// informacion

// crear una estructura con todas las variables que se van a
// remplazar en el template

// crear un metodo encargado de enviar el mail
// cargar el template ParseTemplate("carpeta/archivo.html", estructuraDeVariablesASustituir)

// mailData es de tipo BaseMail
// Enviar el correo con SendMail(mailData)

// SendMail() retorna nil si todo salio bien


// Dataformat for the template

type VerificationCodeMailData struct {
	Host string
	Email string
	Token string
	Username string
}

type PasswordInstructionMailData struct {
	Host string
	Username string
	Code string
	Email string
}

// Send methods
// Metodos de envio

func SendConfirmationCode(user models.User) error{

	// Parse Template
	body, err := ParseTemplate("account/confirmation.html",VerificationCodeMailData{
		Host: os.Getenv("HOST"),
		Email: user.Email,
		Token: user.ConfirmationToken,
		Username: user.FirstName,
	})

	if err != nil {
		log.Fatal(err)
		return err
	} else{

		// Build mail data
		mailData := BaseMail{
			To: []string{user.Email},
			Body: body,
			Subject: "Verificación de cuenta",
			From: "no-contestar@taxky.win",
		}

		// Send Mail
		result := SendMail(mailData)
		return result

	}

}

func SendPasswordInstruction(user models.User) error {

	// Parse Template
	body, err := ParseTemplate("account/password_instructions.html",PasswordInstructionMailData{
		Host: os.Getenv("HOST"),
		Username: user.FirstName,
		Code: user.ResetPasswordToken,
		Email: user.Email,
	})

	if err != nil {
		log.Fatal(err)
		return err
	} else{

		// Build mail data
		mailData := BaseMail{
			To: []string{user.Email},
			Body: body,
			Subject: "Contraseña olvidada",
			From: "no-contestar@taxky.win",
		}

		// Send Mail
		result := SendMail(mailData)
		return result

	}

}