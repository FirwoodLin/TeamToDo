package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
)

// GetTasksByDeadline 获取*截止时间-30min*在startTime和endTime之间的任务列表
func GetTasksByDeadline(startTime string, endTime string) (tasks []model.Task, err error) {
	err = global.Sql.
		Where("deadline BETWEEN ? AND ?", startTime, endTime).
		Where("taskStatus != ?", model.Finished).
		Find(&tasks).Error
	return
}

// GetTasksByStartTime 获取*开始时间*在startTime和endTime之间的任务列表
func GetTasksByStartTime(startTime string, endTime string) (tasks []model.Task, err error) {
	err = global.Sql.
		Where("startAt BETWEEN ? AND ?", startTime, endTime).
		Where("taskStatus = ?", model.NotStarted).
		Find(&tasks).Error
	return
}
