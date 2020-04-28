package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/taxky/models"
	"gopkg.in/validator.v2"
)

type DataLogin struct {
	Email    string
	Password string
}

type DataLoginEmployees struct {
	User     string
	Password string
}

type TokenData struct {
	Token string
}

type RecoveryPasswordData struct {
	Token    string
	Password string
}

func GenerateSecureToken() (string, error) {
	return randomHex(32)
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GetToken(_ctx echo.Context) string {
	/*
		We look for the Authorization key within the headers that
		the petition receives

		Buscamos la llave Authorization dentro de los headers que
		recibe la petición
	*/
	auth := string(_ctx.Request().Header.Get("Authorization"))

	return auth
}

func Authentication(_ctx echo.Context, _DB *gorm.DB) (uint, error) {

	user := &models.User{}
	/*
		We get the token

		Obtenemos el Token
	*/
	token := GetToken(_ctx)

	log.Printf("Token", token)
	/*
		We validate that the token does not come empty

		Validamos que el token no venga vacío
	*/
	if token == "" {
		return 0, errors.New("An authorization header is required")
	}

	if err := _DB.Where("token = ?", token).First(user).Error; err != nil {

		return 0, errors.New("Unauthorized")
	} else {
		return user.ID, nil
	}

	/*
		Buscamos la llave Authorization dentro de los headers que
		recibe la petición

		Retornamos nil para indicar que no se presento ningún error
		en el proceso de validación
	*/
	return 0, nil
}

func Access(_ctx echo.Context, _DB *gorm.DB) (interface{}, error) {
	data_login := &DataLogin{}
	user := &models.User{}

	if err := _ctx.Bind(data_login); err != nil {
		return nil, errors.New("An error has occurred while processing the received data")
	}

	data_login.Password = Encrypt(data_login.Password)

	if err := _DB.Where("email = ? AND password = ?", data_login.Email, data_login.Password).First(&user).Error; err != nil {
		return nil, errors.New("El usuario o contraseña no coincide con ningun registro.")
	}

	return user, nil
}

// Confirma la cuenta
func ConfirmAccount(_ctx echo.Context, _DB *gorm.DB) error {

	user := &models.User{}

	token := &TokenData{}

	if err := _ctx.Bind(token); err != nil {
		return errors.New("an error has occurred while processing the received data")
	}

	err := _DB.Where("confirmation_token = ?", token.Token).First(user).Error
	if err != nil {
		return errors.New("usuario no encontrado, codigo invalido")
	}

	time := time.Now()
	user.ConfirmedAt = &time

	if err := _DB.Save(user).Error; err != nil {
		myerr, ok := err.(*mysql.MySQLError)
		if !ok {
			return err
		}
		if myerr.Number == 1062 {
			return err
		} else {
			return err
		}
	}

	return nil
}

// Actualiza la contrasena por medio del recovery token
func PasswordRecovery(_ctx echo.Context, _DB *gorm.DB) error {

	user := &models.User{}
	data := &RecoveryPasswordData{}

	// Parse body data
	if err := _ctx.Bind(data); err != nil {
		return errors.New("an error has occurred while processing the received data")
	}

	// Find user
	err := _DB.Where("reset_password_token = ?", data.Token).First(user).Error
	if err != nil {
		return errors.New("usuario no encontrado, codigo invalido")
	}

	// Validate password
	if err := validator.Validate(user); err != nil {
		return err
	}

	// Encrypt password
	password := Encrypt(data.Password)
	user.Password = password

	// Update password
	if err := _DB.Save(user).Error; err != nil {
		myerr, ok := err.(*mysql.MySQLError)
		if !ok {
			return err
		}
		if myerr.Number == 1062 {
			return err
		} else {
			return err
		}
	}

	return nil

}
