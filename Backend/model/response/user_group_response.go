package response

import "TeamToDo/model"

type UserGroupResponse struct {
	model.TimeModel
	// 联合主键
	UserID  uint         `gorm:"primaryKey;column:userID"`
	GroupID uint         `gorm:"primaryKey;column:groupID"`
	Role    model.Role   `gorm:"column:role"`
	User    UserResponse `gorm:"foreignKey:userID"`
	Group   model.Group  `gorm:"foreignKey:groupID"`
}
