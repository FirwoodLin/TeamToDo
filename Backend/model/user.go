package model

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	TimeModel
	UserID     uint                `json:"userID" gorm:"primarykey"` // 使用 UserID 而非 UserId
	UserName   string              `json:"userName" gorm:"type:varchar(20)" validate:"max=20"`
	UserAvatar string              `json:"userAvatar" gorm:"type:varchar(100)" validate:"uri"`
	Email      string              `json:"email" gorm:"type:varchar(100);unique" validate:"required,email,max=100"`
	Password   string              `json:"password" gorm:"size:60" validate:"required,len=60"`
	IsVerified bool                `json:"isVerified" gorm:"isVerified"`
	Groups     []UserGroup         `json:"-" gorm:"foreignKey:UserID"`
	validate   *validator.Validate `gorm:"-"`
}
type EmailVerification struct {
	TimeModel
	Email string `gorm:"column:email;type:varchar(100)" validate:"required,email,max=100"`
	Uuid  string `gorm:"unique;column:uuid;not null;type:varchar(255)"`
}

func (u *User) Validate() error {
	if u.validate == nil {
		u.validate = validator.New()
	}
	return u.validate.Struct(u)
}
