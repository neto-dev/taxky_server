package v1

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/taxky/models"
	"github.com/taxky/utils"
	validator "gopkg.in/validator.v2"
	"net/http"
)

type AwardCharacters struct {
	Awards     []models.Award
	Characters []models.Character
}
type ArrayIDs struct {
	IDs []int
}

type ControllerAward struct {
	/*
		We inherited from the base structure
		Heredamos de la estructura base
	*/
	BaseController
	DB *gorm.DB
}

func (_controller_ ControllerAward) Get(_ctx echo.Context) error {
	return Get(_ctx, _controller_.DB, &[]models.Award{})
}

func (_controller_ ControllerAward) GetByID(_ctx echo.Context) error {
	return GetByID(_ctx, _controller_.DB, &models.Award{})
}

func (_controller_ ControllerAward) Filters(_ctx echo.Context) error {
	return Filters(_ctx, _controller_.DB, &[]models.Award{})
}

func (_controller_ ControllerAward) Create(_ctx echo.Context) error {
	return Create(_ctx, _controller_.DB, &models.Award{})
}

func (_controller_ ControllerAward) AwardCharacter(_ctx echo.Context) error {
	//Authenticacion

	var newAwards []uint
	AwardCharacters := &AwardCharacters{}

	_, err := utils.Authentication(_ctx, _controller_.DB)

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	/*
		We recover the parameters we received in the request
		Recuperamos los parametros que recibimos en el request
	*/
	if err := _ctx.Bind(&AwardCharacters); err != nil {
		fmt.Println(err)
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		We validate that the parameters have the correct structure
		Validamos que los parametros tengan la estructura correcta
	*/

	if err := validator.Validate(&AwardCharacters); err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	for i := 0; i < len(AwardCharacters.Awards); i++ {
		data := AwardCharacters.Awards[i]

		/*
			We created the new record in the database
			Creamos el nuevo registro en la base de datos
		*/
		if err := _controller_.DB.Create(&data).Error; err != nil {
			return utils.ReturnErrorJSON(_ctx, "Record Create Failure", err)
		}

		newAwards = append(newAwards, data.ID)

		for i := 0; i < len(AwardCharacters.Characters); i++ {
			character := &models.Character{}
			if err := _controller_.DB.Find(&character, AwardCharacters.Characters[i].ID).Error; err != nil {
				return utils.ReturnErrorJSON(_ctx, "Record not found", err)
			}
			if err := _controller_.DB.Model(&character).Association("Awards").Append(data).Error; err != nil {
				return utils.ReturnErrorJSON(_ctx, "Record Create Failure", err)
			}
		}
	}

	awards := []models.Award{}

	if err := _controller_.DB.Where("id IN (?)", newAwards).Find(&awards).Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, "Records Update Failure", err)
	}

	/*
		We return the values
		Retornamos los valores
	*/
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, awards)
}

func (_controller_ ControllerAward) Update(_ctx echo.Context) error {
	return Update(_ctx, _controller_.DB, &models.Award{})
}

func (_controller_ ControllerAward) Delete(_ctx echo.Context) error {
	return Delete(_ctx, _controller_.DB, &models.Award{})
}

func NewAwardController(_e echo.Group, _DB *gorm.DB) echo.Group {
	/*
		Se declara la variable controller la cual contendra toda la estructura de nuestros controller methods.
		The controller variable is declared which will contain all the structure of our controller methods.
	*/
	controller := &ControllerAward{
		DB: _DB,
	}

	/*
		Se declaran las rutas correspondientes a cada metodo
		The routes corresponding to each method are declared.
	*/

	_e.GET("/awards", controller.Get)
	_e.GET("/awards/:id", controller.GetByID)
	_e.POST("/awards/filters", controller.Filters)
	_e.POST("/awards", controller.Create)
	_e.PUT("/awards/:id", controller.Update)
	_e.PATCH("/awards/:id", controller.Update)
	_e.DELETE("/awards/:id", controller.Delete)

	return _e
}
