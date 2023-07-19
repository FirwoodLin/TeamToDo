package model

type Remark struct {
	TimeModel
	RemarkID  string `json:"remarkID" gorm:"primary_key" ` // 评论ID
	Content   string `json:"content" gorm:"content"`       // 评论内容
	CreatedAt string `json:"createdAt" gorm:"createdAt"  ` // 评论日期
	UpdatedAt string `json:"updatedAt" gorm:"updatedAt"`   // 修改日期
	// 关联
	UserID uint `json:"userID" gorm:"userID"` // 评论者的ID
	User   User `json:"user"`                 // 评论者的信息
	TaskID uint `json:"taskID" gorm:"taskID"` // 任务ID
	Task   Task `json:"task"`                 // 任务信息
}
