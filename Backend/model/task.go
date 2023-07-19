package model

import (
	"time"
)

type TaskStatus int

// 任务状态
const (
	NotStarted TaskStatus = iota
	Processing
	Finished
)

type Task struct {
	GroupID     uint       `json:"groupID"`               // 任务所属的团队的ID
	TaskName    string     `json:"taskName"`              // 任务名称
	Description *string    `json:"description,omitempty"` // 任务描述
	TaskStatus  TaskStatus `json:"taskStatus"`            // 任务状态
	// 时间相关
	Deadline       time.Time `json:"deadline"`       // 任务截止日期
	RemindBefore   int       `json:"remindBefore"`   // 结束前x分钟提醒
	RemindWhen     time.Time `json:"remindWhen"`     // 特定时间提醒
	RemindWhenDone bool      `json:"remindWhenDone"` // 成员完成时提醒
	// 所有者相关
	OwnerID uint `json:"-" gorm:"foreignKey:userID"` // 任务所有者的ID
	Owner   User `json:"owner"`
}
