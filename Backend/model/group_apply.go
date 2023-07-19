package model

type ApplyStatus int

const (
	ApplyStatusPending ApplyStatus = iota
	ApplyStatusAgreed
	ApplyStatusRejected
)

// GroupApply 加群申请(需要管理员审核)
type GroupApply struct {
	TimeModel
	GroupApplyID uint `json:"groupApplyID" gorm:"column:groupApply;primaryKey"`
	GroupID      uint `json:"groupID" gorm:"column:groupID"`
	UserID       uint `json:"userID" gorm:"column:userID"`
	//UserName     string      `json:"userName" gorm:"column:userName"`
	Status ApplyStatus `json:"status" gorm:"column:status"`
	User   User        `json:"user" gorm:"foreignKey:userID"`
}

type GroupJoinCode struct {
	TimeModel
	GroupJoinCodeID uint `json:"groupJoinCodeID" gorm:"column:groupJoinCodeID;primaryKey"`
	// 正文
	Code     string `json:"code" gorm:"column:code;unique"`
	GroupID  uint   `json:"groupID" gorm:"column:groupID"`
	ExpireAt int64  `json:"expireAt" gorm:"column:expireAt"` // 使用 Unix 时间；否则可能出现时区问题
	// 进行外键关联
	Inviter   User
	InviterID uint `json:"inviterID" gorm:"column:inviterID;foreignKey:userID"`
}
