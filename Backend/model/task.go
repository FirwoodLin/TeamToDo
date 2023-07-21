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
	TaskID      uint       `gorm:"primaryKey;column:taskID" json:"taskID"`          // 任务ID
	TaskName    string     `json:"taskName" gorm:"column:taskName"`                 // 任务名称
	Description *string    `json:"description,omitempty" gorm:"column:description"` // 任务描述
	TaskStatus  TaskStatus `json:"taskStatus" gorm:"column:taskStatus"`             // 任务状态
	// 时间相关
	StartAt        time.Time `json:"startAt" gorm:"column:startAt"`               // 任务开始时间
	Deadline       time.Time `json:"deadline" gorm:"column:deadline"`             // 任务截止日期
	RemindBefore   int       `json:"remindBefore" gorm:"column:remindBefore"`     // 结束前x分钟提醒
	RemindWhen     time.Time `json:"remindWhen" gorm:"column:remindWhen"`         // 特定时间提醒
	RemindWhenDone bool      `json:"remindWhenDone" gorm:"column:remindWhenDone"` // 成员完成时提醒
	// 所有者相关
	OwnerID uint  `json:"ownerID" gorm:"column:ownerID"` // 任务所有者的ID
	Owner   User  `json:"owner"`
	GroupID uint  `json:"groupID" gorm:"column:groupID"` // 任务所属的团队的ID
	Group   Group `json:"group" `                        // 任务所属的团队
}

// UserTask 用户-任务关系表；记录用户的任务完成信息 - deprecated
type UserTask struct {
	TimeModel
	// 联合主键
	UserID uint `gorm:"primaryKey;column:userID"`
	TaskID uint `gorm:"primaryKey;column:taskID"`
	// 任务状态
	TaskStatus TaskStatus `gorm:"column:taskStatus"`
}
