package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Award struct {
	gorm.Model
		Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(250)"`
	Points      int32
	UserID      uint
	User        User
	Status      bool
}

func (model *Award) BeforeUpdate() (err error) {
	model.CreatedAt = time.Now()
	return
}
