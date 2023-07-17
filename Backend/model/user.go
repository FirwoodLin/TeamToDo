package model

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	TimeModel
	UserID     uint                `json:"userID" gorm:"primarykey"` // 使用 UserID 而非 UserId
	UserName   string              `json:"userName" gorm:"type:varchar(20);unique" validate:"required,max=20"`
	UserAvatar string              `json:"userAvatar" gorm:"type:varchar(100)" validate:"uri"`
	Email      string              `json:"email" gorm:"type:varchar(100);unique" validate:"required,email,max=100"`
	Password   string              `json:"password" gorm:"size:60" validate:"required,len=60"`
	validate   *validator.Validate `gorm:"-"`
}

func (u *User) Validate() error {
	if u.validate == nil {
		u.validate = validator.New()
	}
	return u.validate.Struct(u)
}
