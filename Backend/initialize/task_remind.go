package initialize

import (
	"TeamToDo/database"
	"TeamToDo/global"
	"TeamToDo/model"
	"github.com/jinzhu/copier"
	"time"
)

var isFirst bool = true

type TaskRemind struct {
	// 任务基本信息
	TaskID      uint             `gorm:"primaryKey;column:taskID" json:"taskID"` // 任务ID
	TaskName    string           `json:"taskName"`                               // 任务名称
	Description *string          `json:"description,omitempty"`                  // 任务描述
	TaskStatus  model.TaskStatus `json:"taskStatus"`                             // 任务状态
	// 时间相关
	Deadline time.Time `json:"deadline"` // 任务截止日期
	GroupID  uint      `json:"groupID"`  // 任务所属的团队的ID
	StartAt  time.Time `json:"startAt"`  // 任务开始时间
	// 提醒类型
	NoticeTime time.Time `json:"noticeTime"` // 提醒时间
}

func Scheduler() {
	for {
		if !isFirst {
			timer := time.NewTicker(1 * time.Hour)
			<-timer.C
		}
		isFirst = false

		// 获取接下来一小时内需要发送的任务列表，并进行合并:(截止日期将近 且 没有完成) & (开始日期将近 且 没有开始)
		// 距离截止日期 30min 的任务
		tasksDDL, err := database.GetTasksByDeadline(formatTime(0), formatTime(1))
		if err != nil {
			global.Logger.Errorf("查询一小时任务列表出错：%v", err)
			return
		}
		// 下一个小时开始的任务
		tasksStart, err := database.GetTasksByStartTime(formatTime(0), formatTime(1))
		//var tasks []model.Task
		//tasks = append(tasks, tasksDDL...)
		//tasks = append(tasks, tasksStart...)
		if err != nil {
			global.Logger.Errorf("查询一小时任务列表出错：%v", err)
			return
		}
		// 任务通道
		ch := make(chan TaskRemind, 100)         // 容量100
		taskFunc := func(ch <-chan TaskRemind) { // 新建处理任务函数: 从通道中读取任务并发送给 remindTask
			for item := range ch {
				go remindTask(item)
			}
		}
		go taskFunc(ch) // 开启处理任务协程
		// 投递任务
		for _, task := range tasksDDL {
			var taskRemind TaskRemind
			_ = copier.Copy(&taskRemind, &task)
			taskRemind.NoticeTime = task.Deadline.Add(-30 * time.Minute) // 设置提醒时间为截止日期前30min
			ch <- taskRemind
		}
		for _, task := range tasksStart {
			var taskRemind TaskRemind
			_ = copier.Copy(&taskRemind, &task)
			taskRemind.NoticeTime = task.StartAt // 设置提醒时间为开始时间
			ch <- taskRemind
		}
	}
}

// 获取当前时间加上 hourToAdd 个小时的时间字符串
func formatTime(hourToAdd time.Duration) string {
	return time.Now().Add(time.Hour * hourToAdd).Format("2006-01-02 15:04:05")
}
func remindTask(task TaskRemind) {
	// 设置发件时间（定时器）
	diff := task.NoticeTime.Sub(time.Now())
	timer := time.NewTimer(diff)
	<-timer.C
	// 获取任务的参与者
	userGroups, err := database.FindGroupMembers(task.GroupID)
	if err != nil {
		global.Logger.Errorf("提醒模块-获取任务参与者出错：%v", err)
		return
	}
	// 发送邮件
	for _, userGroup := range userGroups {
		// 获取参与者邮箱
		user, err := database.UserQueryOneAllInfo(userGroup.UserID)
		if err != nil {
			global.Logger.Errorf("提醒模块-获取参与者邮箱出错：%v", err)
			return
		}
		// 发送邮件 TODO：完成SendMail 函数
		err = SendEmail(user)
		if err != nil {
			global.Logger.Errorf("提醒模块-发送邮件出错：%v", err)
			return
		}
	}
}
func SendEmail(user *model.User) error {
	return nil
}
