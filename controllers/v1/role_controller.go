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
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/taxky/models"

	//validator "gopkg.in/validator.v2"
	"net/http"

	"github.com/taxky/utils"
	validator "gopkg.in/validator.v2"
)

type ControllerRole struct {
	/*
		We inherited from the base structure

		Heredamos de la estructura base
	*/
	BaseController
	DB *gorm.DB
}

type MultipleRoles struct {
	Roles []models.Role
}

func (_controller_ ControllerRole) Get(_ctx echo.Context) error {
	return Get(_ctx, _controller_.DB, &[]models.Role{})
}

func (_controller_ ControllerRole) GetByID(_ctx echo.Context) error {
	return GetByID(_ctx, _controller_.DB, &models.Role{})
}

func (_controller_ ControllerRole) Filters(_ctx echo.Context) error {
	return Filters(_ctx, _controller_.DB, &[]models.Role{})
}

func (_controller_ ControllerRole) Create(_ctx echo.Context) error {
	return Create(_ctx, _controller_.DB, &models.Role{})
}

func (_controller_ ControllerRole) Update(_ctx echo.Context) error {
	return Update(_ctx, _controller_.DB, &models.Role{})
}

func (_controller_ ControllerRole) Delete(_ctx echo.Context) error {
	return Delete(_ctx, _controller_.DB, &models.Role{})
}

func (_controller_ ControllerRole) MultipleActions(_ctx echo.Context) error {
	//Authenticacion

	multipleRoles := &MultipleRoles{}

	_, err := utils.Authentication(_ctx, _controller_.DB)

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	/*
		We recover the parameters we received in the request

		Recuperamos los parametros que recibimos en el request
	*/
	if err := _ctx.Bind(&multipleRoles); err != nil {
		fmt.Println(err)
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		We validate that the parameters have the correct structure

		Validamos que los parametros tengan la estructura correcta
	*/

	if err := validator.Validate(&multipleRoles); err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	for i := 0; i < len(multipleRoles.Roles); i++ {
		role := &models.Role{}

		roleValues := multipleRoles.Roles[i]

		role.ID = roleValues.ID

		/*
			We created the new record in the database

			Creamos el nuevo registro en la base de datos
		*/
		if role.ID == 0 {
			role.Name = roleValues.Name
			role.Quota = roleValues.Quota
			role.Status = roleValues.Status
			if err := _controller_.DB.Create(&role).Error; err != nil {
				return utils.ReturnErrorJSON(_ctx, "Record Create Failure", err)
			}
		} else {
			if err := _controller_.DB.First(&role, role.ID).Error; err != nil {
				return utils.ReturnErrorJSON(_ctx, "Record Find Failure", err)
			}
			role.Name = roleValues.Name
			role.Quota = roleValues.Quota
			role.Status = roleValues.Status
			if err := _controller_.DB.Save(&role).Error; err != nil {
				return utils.ReturnErrorJSON(_ctx, "Record Create Failure", err)
			}
		}
		/*
			We return the values

			Retornamos los valores
		*/
	}
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, multipleRoles)
}

func NewRoleController(_e echo.Group, _DB *gorm.DB) echo.Group {
	/*
		Se declara la variable controller la cual contendra toda la estructura de nuestros controller methods.

		The controller variable is declared which will contain all the structure of our controller methods.
	*/
	controller := &ControllerRole{
		DB: _DB,
	}

	/*
		Se declaran las rutas correspondientes a cada metodo

		The routes corresponding to each method are declared.
	*/

	_e.GET("/roles", controller.Get)
	_e.GET("/roles/:id", controller.GetByID)
	_e.POST("/roles/filters", controller.Filters)
	_e.POST("/roles", controller.Create)
	_e.POST("/roles/multiple", controller.MultipleActions)
	_e.PUT("/roles/:id", controller.Update)
	_e.PATCH("/roles/:id", controller.Update)
	_e.DELETE("/roles/:id", controller.Delete)

	return _e
}
