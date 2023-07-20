package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"TeamToDo/utils"
)

// TaskCreate 新建任务(需要检查用户身份后调用）返回值直接从传入的 task 读取
func TaskCreate(task *model.Task) (err error) {
	err = global.Sql.Create(task).Error
	if err != nil {
		global.Logger.Infof("创建任务: owner:%v,groupID:%v,错误：%v\n", task.OwnerID, task.GroupID, err)
		return err
	}
	global.Logger.Debug("创建任务: ", *task)
	return nil
}

// TaskUpdate 更新任务
func TaskUpdate(task *model.Task) error {
	//var task model.Task
	err := global.Sql.Model(&model.Task{}).
		Where("taskID = ?", task.TaskID).
		Updates(*task).
		Error
	if err != nil {
		global.Logger.Infof("更新任务状态失败,taskID: %v,err:%v", task.TaskID, err.Error())
		return err
	}
	global.Logger.Debug("更新任务状态: ", task)
	return nil
}

// DeleteTask 删除任务
func DeleteTask(taskID uint) error {
	err := global.Sql.Where("taskID = ?", taskID).Delete(&model.Task{}).Error
	if err != nil {
		global.Logger.Infof("删除任务失败,taskID: %v,err:%v", taskID, err.Error())
		return err
	}
	global.Logger.Debug("删除任务: ", taskID)
	return nil
}

// QueryTasksByGroupID 根据群组 ID 查询任务
func QueryTasksByGroupID(groupID uint) (tasks []model.Task, err error) {
	err = global.Sql.Where("groupID = ?", groupID).Find(&tasks).Error
	if err != nil {
		global.Logger.Infof("查询任务失败,groupID: %v,err:%v", groupID, err.Error())
		return nil, err
	}
	global.Logger.Debug("groupID", groupID, "\ntasks: ", tasks)
	return tasks, err
}

// QueryTasksByGroupIDs 根据 多个群组 ID 查询任务
func QueryTasksByGroupIDs(groupIDs []uint) (tasks []model.Task, err error) {
	err = global.Sql.Where("groupID IN ?", groupIDs).Find(&tasks).Error
	global.Logger.Debug("groupIDs", groupIDs, "\ntasks: ", tasks)
	return
}

// GetTaskIdsFromUserTasks 工具函数，从 UserTask 切片中获取 TaskID 切片
func GetTaskIdsFromUserTasks(userTasks []model.UserTask) (taskIDs []uint) {
	for _, userTask := range userTasks {
		taskIDs = append(taskIDs, userTask.TaskID)
	}
	global.Logger.Debug("userTasks: ", userTasks, "\ntaskIDs: ", taskIDs)
	return taskIDs
}

// QueryTasks 根据 用户群组[]uint,群组 []uint, 关键词 string 查询任务
func QueryTasks(userGroupIDs []uint, groupIDs []uint, keyword string) (tasks []model.Task, err error) {
	db := global.Sql.Model(&model.Task{})
	var intersect []uint
	if len(userGroupIDs) != 0 {
		//db = db.Where("groupID IN ?", userGroupIDs)
		intersect = userGroupIDs
	}
	if len(groupIDs) != 0 {
		//db = db.Where("groupID IN ?", groupIDs)
		if len(intersect) != 0 {
			intersect = utils.Intersect(intersect, groupIDs)
		} else {
			intersect = groupIDs
		}
	}
	db = db.Where("groupID IN ?", intersect)

	if keyword != "" {
		db = db.Where("taskName LIKE ?", "%"+keyword+"%")
	}
	err = db.Find(&tasks).Error
	if err != nil {
		global.Logger.Infof("查询任务失败,groupIDs: %v,err:%v", groupIDs, err.Error())
		return nil, err
	}
	global.Logger.Debug("查询任务 len(tasks): ", len(tasks))
	return tasks, err
}

// QueryTaskByTaskID 根据任务 ID 查询任务 - 保留，作为基本函数
func QueryTaskByTaskID(taskID uint) (task model.Task, err error) {
	err = global.Sql.Where("taskID = ?", taskID).First(&task).Error
	if err != nil {
		global.Logger.Infof("查询任务失败,taskID: %v,err:%v", taskID, err.Error())
		return task, err
	}
	global.Logger.Debug("taskID", taskID, "\ntask: ", task)
	return task, err
}
