package request

import "time"

// TaskCreateRequest 创建任务请求
// OwnerID 根据 token 获取
type TaskCreateRequest struct {
	// 所有者相关
	OwnerID uint `json:"-"`       // 任务所有者的ID
	GroupID uint `json:"groupID"` // 任务所属的团队的ID
	// 任务信息
	TaskName    string  `json:"taskName"`              // 任务名称
	Description *string `json:"description,omitempty"` // 任务描述
	// 时间相关
	Deadline       time.Time `json:"deadline"`       // 任务截止日期
	RemindBefore   int       `json:"remindBefore"`   // 结束前x分钟提醒
	RemindWhen     time.Time `json:"remindWhen"`     // 特定时间提醒
	RemindWhenDone bool      `json:"remindWhenDone"` // 成员完成时提醒
}
