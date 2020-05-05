/*
 * Ernesto Valenzuela Vargas. Internal License
 *
 * Contact: neto.dev@protonmail.com
 *
 (License Content)
*/

/*
 * Revision History:
 *     Initial:        2018/08/24        Ernesto Valenzuela V
 */

package v1

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/taxky/mailers"
	"github.com/taxky/models"
	"github.com/taxky/utils"
	"gopkg.in/validator.v2"
)

type ControllerUser struct {
	/*
		We inherited from the base structure

		Heredamos de la estructura base
	*/
	BaseController
	DB *gorm.DB
}

type AccountData struct {
	Email string
}

func (_controller_ ControllerUser) Login(_ctx echo.Context) error {
	user, err := utils.Access(_ctx, _controller_.DB)

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	return utils.ReturnRespondJSON(_ctx, http.StatusOK, user)
}

func (_controller_ ControllerUser) Get(_ctx echo.Context) error {
	return Get(_ctx, _controller_.DB, &[]models.User{})
}

func (_controller_ ControllerUser) GetByID(_ctx echo.Context) error {
	user := &models.User{}
	/*
		Filter within the database searching for the record corresponding
		to the received id. The preload works to bring up the relationship
		data
		Filtramos dentro de la base de datos buscando el registro que
		corresponda con el id recibido. El preload funciona para traer
		los datos de la relacion
	*/

	idP, err := strconv.Atoi(_ctx.Param("id"))

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		Seteamos la variable idP en la variable id transformandola en uint
		We set the variable idP in the variable id transforming it into uint
	*/
	id := uint(idP)

	if err := _controller_.DB.Preload("Characters").First(user, id).Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, "Record not found", err)
	}

	/*
		We return the values
		Retornamos los valores
	*/
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, user)
}

func (_controller_ ControllerUser) Filters(_ctx echo.Context) error {
	return Filters(_ctx, _controller_.DB, &[]models.User{})
}

func (_controller_ ControllerUser) Create(_ctx echo.Context) error {
	user := &models.User{}

	/*
		We recover the parameters we received in the request

		Recuperamos los parametros que recibimos en el request
	*/
	if err := _ctx.Bind(user); err != nil {
		fmt.Println(err)
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		We validate that the parameters have the correct structure

		Validamos que los parametros tengan la estructura correcta
	*/

	if err := validator.Validate(user); err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		We created the new record in the database

		Creamos el nuevo registro en la base de datos
	*/

	token := utils.GenerateToken(user.Email, user.Password)

	user.Token = token

	password := utils.Encrypt(user.Password)

	user.Password = password

	if err := _controller_.DB.Create(user).Error; err != nil {
		myerr, ok := err.(*mysql.MySQLError)
		if !ok {
			return utils.ReturnErrorJSON(_ctx, "Record Create Failure", err)
		}
		if myerr.Number == 1062 {
			return utils.ReturnErrorJSON(_ctx, "Duplicate Record", err)
		} else {
			return utils.ReturnErrorJSON(_ctx, "Record Create Failure", err)
		}
	}
	/*
		We return the values

		Retornamos los valores
	*/

	return utils.ReturnRespondJSON(_ctx, http.StatusOK, user)
}

func (_controller_ ControllerUser) Update(_ctx echo.Context) error {
	return Update(_ctx, _controller_.DB, &models.User{})
}

func (_controller_ ControllerUser) Delete(_ctx echo.Context) error {
	return Delete(_ctx, _controller_.DB, &models.User{})
}

func (_controller_ ControllerUser) ConfirmAccount(_ctx echo.Context) error {
	err := utils.ConfirmAccount(_ctx, _controller_.DB)

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	return utils.ReturnRespondJSON(_ctx, http.StatusOK, nil)
}

func (_controller_ ControllerUser) PasswordRecovery(_ctx echo.Context) error {
	err := utils.PasswordRecovery(_ctx, _controller_.DB)

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	return utils.ReturnRespondJSON(_ctx, http.StatusOK, nil)
}

func (_controller_ ControllerUser) SendConfirmationEmail(_ctx echo.Context) (err error) {

	user := &models.User{}

	// Parse body
	account := &AccountData{}
	if err := _ctx.Bind(account); err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), errors.New("an error has occurred while processing the received data"))
	}

	//
	// Find User

	err = _controller_.DB.Where("email = ?", account.Email).First(user).Error

	if err == nil {

	GenerateToken:

		// Generate & assign token
		token, err := utils.GenerateSecureToken()

		tmpuser := &models.User{}
		if !_controller_.DB.Where("confirmation_token = ?", token).First(tmpuser).RecordNotFound() {
			log.Println("token duplicado, generando uno nuevo")
			goto GenerateToken
		}

		user.ConfirmationToken = token

		if err == nil {

			log.Println(user.ConfirmationToken)

			if err := _controller_.DB.Save(user).Error; err != nil {
				myerr, ok := err.(*mysql.MySQLError)
				if !ok {
					return utils.ReturnErrorJSON(_ctx, err.Error(), err)
				}
				if myerr.Number == 1062 {
					return utils.ReturnErrorJSON(_ctx, err.Error(), err)
				} else {
					return utils.ReturnErrorJSON(_ctx, err.Error(), err)
				}
			} else {
				// Send email
				err = mailers.SendConfirmationCode(*user)

			}

		} else {
			return utils.ReturnErrorJSON(_ctx, err.Error(), err)
		}

	}

	return utils.ReturnRespondJSON(_ctx, http.StatusOK, err)

}

func (_controller_ ControllerUser) SendPasswordRecoveryEmail(_ctx echo.Context) (err error) {

	user := &models.User{}

	// Parse body
	account := &AccountData{}
	if err := _ctx.Bind(account); err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), errors.New("an error has occurred while processing the received data"))
	}

	// Find User
	err = _controller_.DB.Where("email = ?", account.Email).First(user).Error

	if err == nil {

	GenerateToken:

		// Generate & assign token
		token, err := utils.GenerateSecureToken()

		tmpuser := &models.User{}
		if !_controller_.DB.Where("reset_password_token = ?", token).First(tmpuser).RecordNotFound() {
			log.Println("token duplicado, generando uno nuevo")
			goto GenerateToken
		}

		user.ResetPasswordToken = token

		if err == nil {

			if err := _controller_.DB.Save(user).Error; err != nil {
				myerr, ok := err.(*mysql.MySQLError)
				if !ok {
					return utils.ReturnErrorJSON(_ctx, err.Error(), err)
				}
				if myerr.Number == 1062 {
					return utils.ReturnErrorJSON(_ctx, err.Error(), err)
				} else {
					return utils.ReturnErrorJSON(_ctx, err.Error(), err)
				}
			} else {
				// Send email
				err = mailers.SendPasswordInstruction(*user)
			}

		} else {
			return utils.ReturnErrorJSON(_ctx, err.Error(), err)
		}

	}

	return utils.ReturnRespondJSON(_ctx, http.StatusOK, err)

}

func NewUserController(_e echo.Group, _DB *gorm.DB) echo.Group {
	/*
		Se declara la variable controller la cual contendra toda la estructura de nuestros controller methods.

		The controller variable is declared which will contain all the structure of our controller methods.
	*/
	controller := &ControllerUser{
		DB: _DB,
	}

	/*
		Se declaran las rutas correspondientes a cada metodo

		The routes corresponding to each method are declared.
	*/

	_e.POST("/users/login", controller.Login)
	_e.GET("/users", controller.Get)
	_e.GET("/users/:id", controller.GetByID)
	_e.POST("/users/filters", controller.Filters)
	_e.POST("/users", controller.Create)
	_e.PUT("/users/:id", controller.Update)
	_e.PATCH("/users/:id", controller.Update)
	_e.DELETE("/users/:id", controller.Delete)

	_e.POST("/users/account-confirm", controller.ConfirmAccount)
	_e.POST("/users/change-password", controller.PasswordRecovery)

	_e.POST("/users/send-account-confirm", controller.SendConfirmationEmail)
	_e.POST("/users/send-password-forgot", controller.SendPasswordRecoveryEmail)

	return _e
}
