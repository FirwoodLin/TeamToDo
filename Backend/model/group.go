package model

type Role int

// 定义角色的方式
const (
	RoleVisitor Role = iota //  RoleVisitor 不在群组中
	RoleOwner
	RoleAdmin
	RoleMember
)

type Group struct {
	TimeModel
	GroupID     uint        `json:"groupID" gorm:"column:groupID;primaryKey"`
	Description string      `json:"description" gorm:"column:description"`
	GroupName   string      `json:"groupName"  gorm:"column:groupName"` // 群组名称
	GroupAvatar string      `json:"groupAvatar" gorm:"column:groupAvatar"`
	MemberCount int         `json:"memberCount" gorm:"column:description"`
	Members     []UserGroup `gorm:"foreignKey:groupID"`
}

// UserGroup 用户-群组关系表；记录用户加入的群组
type UserGroup struct {
	TimeModel
	// 联合主键
	UserID  int   `gorm:"primaryKey"`
	GroupID int   `gorm:"primaryKey"`
	Role    Role  `gorm:"column:role"`
	User    User  `gorm:"foreignKey:userID"`
	Group   Group `gorm:"foreignKey:groupID"`
}
type ApplyStatus int

const (
	ApplyStatusPending ApplyStatus = iota
	ApplyStatusAgreed
	ApplyStatusRejected
)

type GroupApplication struct {
	TimeModel
	GroupApplicationID uint        `json:"groupApplicationID" gorm:"column:groupApplicationID;primaryKey"`
	GroupID            uint        `json:"groupID" gorm:"column:groupID"`
	UserID             uint        `json:"userID" gorm:"column:userID"`
	UserName           string      `json:"userName" gorm:"column:userName"`
	Status             ApplyStatus `json:"status" gorm:"column:status"`
}
