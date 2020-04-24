package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Permission struct {
	gorm.Model
	Name string `gorm:"type:varchar(200)"`
}

func (model *Permission) BeforeUpdate() (err error) {

	model.CreatedAt = time.Now()
	return
}
