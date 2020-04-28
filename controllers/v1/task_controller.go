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

type TaskCharacters struct {
	Tasks      []models.Task
	Characters []models.Character
}

type ControllerTask struct {
	/*
		We inherited from the base structure
		Heredamos de la estructura base
	*/
	BaseController
	DB *gorm.DB
}

func (_controller_ ControllerTask) Get(_ctx echo.Context) error {
	return Get(_ctx, _controller_.DB, &[]models.Task{})
}

func (_controller_ ControllerTask) GetByID(_ctx echo.Context) error {
	return GetByID(_ctx, _controller_.DB, &models.Task{})
}

func (_controller_ ControllerTask) Filters(_ctx echo.Context) error {
	return Filters(_ctx, _controller_.DB, &[]models.Task{})
}

func (_controller_ ControllerTask) Create(_ctx echo.Context) error {
	//Authenticacion

	var newtasks []uint
	TaskCharacters := &TaskCharacters{}

	_, err := utils.Authentication(_ctx, _controller_.DB)

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	/*
		We recover the parameters we received in the request
		Recuperamos los parametros que recibimos en el request
	*/
	if err := _ctx.Bind(&TaskCharacters); err != nil {
		fmt.Println(err)
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		We validate that the parameters have the correct structure
		Validamos que los parametros tengan la estructura correcta
	*/

	if err := validator.Validate(&TaskCharacters); err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	for i := 0; i < len(TaskCharacters.Tasks); i++ {
		data := TaskCharacters.Tasks[i]

		/*
			We created the new record in the database
			Creamos el nuevo registro en la base de datos
		*/
		if err := _controller_.DB.Create(&data).Error; err != nil {
			return utils.ReturnErrorJSON(_ctx, "Record Create Failure", err)
		}

		newtasks = append(newtasks, data.ID)

		for i := 0; i < len(TaskCharacters.Characters); i++ {
			character := &models.Character{}
			if err := _controller_.DB.Find(&character, TaskCharacters.Characters[i].ID).Error; err != nil {
				return utils.ReturnErrorJSON(_ctx, "Record not found", err)
			}
			if err := _controller_.DB.Model(&character).Association("tasks").Append(data).Error; err != nil {
				return utils.ReturnErrorJSON(_ctx, "Record Create Failure", err)
			}
		}
	}

	tasks := []models.Task{}

	if err := _controller_.DB.Where("id IN (?)", newtasks).Find(&tasks).Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, "Records Update Failure", err)
	}

	/*
		We return the values
		Retornamos los valores
	*/
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, tasks)
}

func (_controller_ ControllerTask) Update(_ctx echo.Context) error {
	return Update(_ctx, _controller_.DB, &models.Task{})
}

func (_controller_ ControllerTask) Delete(_ctx echo.Context) error {
	return Delete(_ctx, _controller_.DB, &models.Task{})
}

func NewTaskController(_e echo.Group, _DB *gorm.DB) echo.Group {
	/*
		Se declara la variable controller la cual contendra toda la estructura de nuestros controller methods.
		The controller variable is declared which will contain all the structure of our controller methods.
	*/
	controller := &ControllerTask{
		DB: _DB,
	}

	/*
		Se declaran las rutas correspondientes a cada metodo
		The routes corresponding to each method are declared.
	*/

	_e.GET("/tasks", controller.Get)
	_e.GET("/tasks/:id", controller.GetByID)
	_e.POST("/tasks/filters", controller.Filters)
	_e.POST("/tasks", controller.Create)
	_e.PUT("/tasks/:id", controller.Update)
	_e.PATCH("/tasks/:id", controller.Update)
	_e.DELETE("/tasks/:id", controller.Delete)

	return _e
}
