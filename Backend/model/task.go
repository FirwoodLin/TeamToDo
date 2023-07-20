package model

import (
	"time"
)

type TaskStatus int

// 任务状态
const (
	NotStarted TaskStatus = iota // 没有开始
	Processing                   // 进行中（已经接收到信息）
	Finished                     // 已完成
	//Deleted                      // 已删除
)

type Task struct {
	TimeModel // 创建/修改时间
	// 任务基本信息
	TaskID      uint       `gorm:"primaryKey;column:taskID" json:"taskID"` // 任务ID
	TaskName    string     `json:"taskName"`                               // 任务名称
	Description *string    `json:"description,omitempty"`                  // 任务描述
	TaskStatus  TaskStatus `json:"taskStatus"`                             // 任务状态
	// 时间相关
	StartAt        time.Time `json:"startAt"`        // 任务开始日期
	Deadline       time.Time `json:"deadline"`       // 任务截止日期
	RemindBefore   int       `json:"remindBefore"`   // 结束前x分钟提醒
	RemindWhen     time.Time `json:"remindWhen"`     // 特定时间提醒
	RemindWhenDone bool      `json:"remindWhenDone"` // 成员完成时提醒
	// 所有者相关
	OwnerID uint  `json:"-" gorm:"foreignKey:userID"` // 任务所有者的ID
	Owner   User  `json:"owner" `
	GroupID uint  `json:"-" gorm:"foreignKey:userID"`      // 任务所属的团队的ID
	Group   Group `json:"group" gorm:"foreignKey:groupID"` // 任务所属的团队
}

// UserTask 用户-任务关系表；记录用户的任务完成信息
type UserTask struct {
	TimeModel
	// 联合主键
	UserID uint `gorm:"primaryKey;column:userID"`
	TaskID uint `gorm:"primaryKey;column:taskID"`
	// 任务状态
	TaskStatus TaskStatus `gorm:"column:taskStatus"`
}
