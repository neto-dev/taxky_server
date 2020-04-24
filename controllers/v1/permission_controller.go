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
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/taxky/models"
	//validator "gopkg.in/validator.v2"
)

type ControllerPermission struct {
	/*
		We inherited from the base structure

		Heredamos de la estructura base
	*/
	BaseController
	DB *gorm.DB
}

func (_controller_ ControllerPermission) Get(_ctx echo.Context) error {
	return Get(_ctx, _controller_.DB, &[]models.Permission{})
}

func (_controller_ ControllerPermission) GetByID(_ctx echo.Context) error {
	return GetByID(_ctx, _controller_.DB, &models.Permission{})
}

func (_controller_ ControllerPermission) Filters(_ctx echo.Context) error {
	return Filters(_ctx, _controller_.DB, &[]models.Permission{})
}

func (_controller_ ControllerPermission) Create(_ctx echo.Context) error {
	return Create(_ctx, _controller_.DB, &models.Permission{})
}

func (_controller_ ControllerPermission) Update(_ctx echo.Context) error {
	return Update(_ctx, _controller_.DB, &models.Permission{})
}

func (_controller_ ControllerPermission) Delete(_ctx echo.Context) error {
	return Delete(_ctx, _controller_.DB, &models.Permission{})
}

func NewPermissionController(_e echo.Group, _DB *gorm.DB) echo.Group {
	/*
		Se declara la variable controller la cual contendra toda la estructura de nuestros controller methods.

		The controller variable is declared which will contain all the structure of our controller methods.
	*/
	controller := &ControllerPermission{
		DB: _DB,
	}

	/*
		Se declaran las rutas correspondientes a cada metodo

		The routes corresponding to each method are declared.
	*/

	_e.GET("/permissions", controller.Get)
	_e.GET("/permissions/:id", controller.GetByID)
	_e.POST("/permissions/filters", controller.Filters)
	_e.POST("/permissions", controller.Create)
	_e.PUT("/permissions/:id", controller.Update)
	_e.PATCH("/permissions/:id", controller.Update)
	_e.DELETE("/permissions/:id", controller.Delete)

	return _e
}
