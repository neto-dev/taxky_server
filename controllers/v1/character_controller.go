package v1

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/taxky/models"
	"github.com/taxky/utils"
	"net/http"
	"strconv"
)

type ControllerCharacter struct {
	/*
		We inherited from the base structure
		Heredamos de la estructura base
	*/
	BaseController
	DB *gorm.DB
}

func (_controller_ ControllerCharacter) Get(_ctx echo.Context) error {
	return Get(_ctx, _controller_.DB, &[]models.Character{})
}

func (_controller_ ControllerCharacter) GetByID(_ctx echo.Context) error {

	character := &models.Character{}
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

	if err := _controller_.DB.Preload("Tasks").Preload("Awards").First(character, id).Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, "Record not found", err)
	}

	/*
		We return the values
		Retornamos los valores
	*/
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, character)
}

func (_controller_ ControllerCharacter) Filters(_ctx echo.Context) error {
	return Filters(_ctx, _controller_.DB, &[]models.Character{})
}

func (_controller_ ControllerCharacter) Create(_ctx echo.Context) error {
	return Create(_ctx, _controller_.DB, &models.Character{})
}

func (_controller_ ControllerCharacter) Update(_ctx echo.Context) error {
	return Update(_ctx, _controller_.DB, &models.Character{})
}

func (_controller_ ControllerCharacter) RemoveAwards(_ctx echo.Context) error {

	character := &models.Character{}

	//Authenticacion

	_, err := utils.Authentication(_ctx, _controller_.DB)

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	idP, err := strconv.Atoi(_ctx.Param("id"))

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		Seteamos la variable idP en la variable id transformandola en uint
		We set the variable idP in the variable id transforming it into uint
	*/
	id := uint(idP)

	/*
		Recover the record you wish to edit
		Recuperamos el registro que se desea editar
	*/
	if err := _controller_.DB.First(character, id).Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, "Record Find Failure", err)
	}

	if err := _controller_.DB.Model(&character).Association("Awards").Clear().Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, "Record Update Failure", err)
	}

	/*
		We return the values
		Retornamos los valores
	*/

	return utils.ReturnRespondJSON(_ctx, http.StatusOK, character)
}

func (_controller_ ControllerCharacter) Delete(_ctx echo.Context) error {
	return Delete(_ctx, _controller_.DB, &models.Character{})
}

func NewCharacterController(_e echo.Group, _DB *gorm.DB) echo.Group {
	/*
		Se declara la variable controller la cual contendra toda la estructura de nuestros controller methods.
		The controller variable is declared which will contain all the structure of our controller methods.
	*/
	controller := &ControllerCharacter{
		DB: _DB,
	}

	/*
		Se declaran las rutas correspondientes a cada metodo
		The routes corresponding to each method are declared.
	*/

	_e.GET("/characters", controller.Get)
	_e.GET("/characters/:id", controller.GetByID)
	_e.POST("/characters/filters", controller.Filters)
	_e.POST("/characters", controller.Create)
	_e.PUT("/characters/:id", controller.Update)
	_e.PATCH("/characters/:id", controller.Update)
	_e.DELETE("/characters/awards/:id", controller.RemoveAwards)
	_e.DELETE("/characters/:id", controller.Delete)

	return _e
}
