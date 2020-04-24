package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Role struct {
	gorm.Model
	Name        string `gorm:"type:varchar(200)"`
	Quota       float32
	Permissions []Permission `gorm:"many2many:roles_permissions;"`
	Status      bool
}

func (model *Role) BeforeUpdate() (err error) {

	model.CreatedAt = time.Now()
	return
}
