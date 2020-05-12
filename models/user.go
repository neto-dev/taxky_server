package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Email              string `gorm:"type:varchar(200)"`
	Password           string `gorm:"type:varchar(250)"`
	ResetPasswordToken string `gorm:"type:varchar(250)"`
	ResetPasswordSend  *time.Time
	SinginCount        int64
	ConfirmationToken  string `gorm:"type:varchar(250)"`
	ConfirmedAt        *time.Time
	FirstName          string `gorm:"type:varchar(250)"`
	LastName           string `gorm:"type:varchar(250)"`
	Status             bool
	Token              string `gorm:"type:varchar(250)"`
	Characters         []Character
	Awards             []Award
	Tasks              []Task
}

func (model *User) BeforeUpdate() (err error) {
	model.CreatedAt = time.Now()
	return
}
