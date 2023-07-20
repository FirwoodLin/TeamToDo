package response
import (
	"TeamToDo/model"
	"time"
)

type TaskResponse struct {
	model.TimeModel // 创建/修改时间
	// 任务基本信息
	TaskID      uint             `gorm:"primaryKey;column:taskID" json:"taskID"` // 任务ID
	TaskName    string           `json:"taskName"`                               // 任务名称
	Description *string          `json:"description,omitempty"`                  // 任务描述
	TaskStatus  model.TaskStatus `json:"taskStatus"`                             // 任务状态
	// 时间相关
	Deadline       time.Time `json:"deadline"`       // 任务截止日期
	RemindBefore   int       `json:"remindBefore"`   // 结束前x分钟提醒
	RemindWhen     time.Time `json:"remindWhen"`     // 特定时间提醒
	RemindWhenDone bool      `json:"remindWhenDone"` // 成员完成时提醒
	// 所有者相关
	OwnerID uint         `json:"-" gorm:"foreignKey:userID"` // 任务所有者的ID
	Owner   UserResponse `json:"owner"`
	GroupID uint         `json:"-"`                               // 任务所属的团队的ID
	Group   model.Group  `json:"group" gorm:"foreignKey:groupID"` // 任务所属的团队

}