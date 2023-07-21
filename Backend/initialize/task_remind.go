package initialize

import (
	"TeamToDo/database"
	"TeamToDo/global"
	"TeamToDo/model"
	"TeamToDo/utils"
	"github.com/jinzhu/copier"
	"time"
)

var isFirst bool = true

const (
	interval = 1000 * time.Second // 每隔 interval 时间执行一次
	before   = 10 * time.Second   // 任务结束前多久提醒
)

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
			timer := time.NewTicker(interval) // 每隔 interval 时间执行一次
			<-timer.C
		}
		isFirst = false

		// 获取接下来 interval 内需要发送的任务列表：(截止日期将近 且 没有完成) ， (开始日期将近 且 没有开始)
		// 距离截止日期 before 的任务
		tasksDDL, err := database.GetTasksByDeadline(formatTime(0), formatTime(interval))
		if err != nil {
			global.Logger.Errorf("查询一小时任务列表出错：%v", err)
			return
		}
		// 下一个 interval 内开始的任务
		tasksStart, err := database.GetTasksByStartTime(formatTime(0), formatTime(interval))
		if err != nil {
			global.Logger.Errorf("查询一小时任务列表出错：%v", err)
			return
		}
		global.Logger.Debugf("DB:最近%v秒任务列表长度：%v", interval.Seconds(), len(tasksDDL)+len(tasksStart))
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
			// 截止时间
			var taskRemind TaskRemind
			_ = copier.Copy(&taskRemind, &task)
			taskRemind.NoticeTime = task.Deadline.Add(-before) // 设置提醒时间为截止时间前 before
			ch <- taskRemind
			global.Logger.Debugf("ch len %v;tasksDDL ID:%v", len(ch), task.TaskID)
		}
		for _, task := range tasksStart {
			// 开始时间
			var taskRemind TaskRemind
			_ = copier.Copy(&taskRemind, &task)
			taskRemind.NoticeTime = task.StartAt // 设置提醒时间为开始时间
			ch <- taskRemind
			global.Logger.Debugf("ch len %v;tasksStart ID:%v", len(ch), task.TaskID)

		}
		close(ch) // 关闭通道
	}
}

// 获取当前时间加上 durationToAdd 时长的时间字符串
func formatTime(durationToAdd time.Duration) string {
	return time.Now().Add(durationToAdd).Add(-8 * time.Hour).Format("2006-01-02 15:04:05")
}

func remindTask(task TaskRemind) {
	// 设置发件时间（定时器）
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now() // 获取当前北京时间 // walkAround
	//now := time.Now().In(loc).Add(time.Hour * 8) // 获取当前北京时间 // walkAround
	diff := task.NoticeTime.Sub(now)
	timer := time.NewTimer(diff)
	global.Logger.Debugf("提醒模块-提醒时间：%v,diff:%v", task.NoticeTime, diff)
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
		// 发送邮件
		err = utils.PostEmail(
			user.Email,
			utils.GenerateRemindMail(task.TaskName,
				*task.Description,
				task.StartAt.Format("2006-01-02 15:04:05"),
				task.Deadline.Format("2006-01-02 15:04:05")))
		if err != nil {
			global.Logger.Errorf("提醒模块-发送邮件出错：%v", err)
			return
		}
	}
}
