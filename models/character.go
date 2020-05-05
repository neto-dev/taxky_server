package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Character struct {
	gorm.Model
	Name     string `gorm:"type:varchar(250)"`
	NickName string `gorm:"type:varchar(250)"`
	Points   int32
	Sex      string `gorm:"type:varchar(10)"`
	Status   bool
	UserID   uint
	User     User
	Tasks    []Task  `gorm:"many2many:characters_tasks;"`
	Awards   []Award `gorm:"many2many:characters_awards;"`
}

func (model *Character) BeforeUpdate() (err error) {
	model.CreatedAt = time.Now()
	return
}
