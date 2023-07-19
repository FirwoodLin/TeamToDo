package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
)

// CreateRemark 发表评论
func CreateRemark(remark *model.Remark) (err error) {
	err = global.Sql.Create(remark).Error
	if err != nil {
		global.Logger.Infof("创建评论: %v,错误：%v\n", remark, err)
	}
	global.Logger.Debug("创建评论: ", *remark)
	return
}

// QueryRemarksByTaskID 查询任务的所有评论；根据任务 ID 查询评论
func QueryRemarksByTaskID(taskID uint) (remarks []model.Remark, err error) {
	err = global.Sql.Model(&model.Remark{}).
		Where("taskID = ?", taskID).
		Find(&remarks).
		Error
	if err != nil {
		global.Logger.Infof("查询任务的所有评论失败,taskID: %v,err:%v", taskID, err.Error())
		return nil, err
	}
	global.Logger.Debug("taskID", taskID, "\nremarks: ", remarks)
	return remarks, err
}

// QueryRemarkByRemarkID 查询评论信息-根据评论ID；鉴权用
func QueryRemarkByRemarkID(remarkID uint) (remark model.Remark, err error) {
	err = global.Sql.
		Where("remarkID = ?", remarkID).
		Find(&remark).Error
	if err != nil {
		global.Logger.Infof("查询评论失败,remarkID: %v,err:%v", remarkID, err.Error())
		return remark, err
	}
	global.Logger.Debug("remarkID", remarkID, "\nremark: ", remark)
	return remark, err
}

// UpdateRemark 编辑评论;将 RemarkID 和 Remark Content 内容存在 remark 参数中
func UpdateRemark(remark *model.Remark) error {
	err := global.Sql.Model(&model.Remark{}).
		Where("remarkID = ?", remark.RemarkID).
		Updates(*remark).
		Error
	if err != nil {
		global.Logger.Infof("更新评论失败,remark: %v,err:%v", remark, err.Error())
		return err
	}
	global.Logger.Debug("更新评论: ", remark)
	return nil
}

// DeleteRemark 删除评论
func DeleteRemark(remarkID uint) error {
	err := global.Sql.Delete(&model.Remark{}, remarkID).Error
	if err != nil {
		global.Logger.Infof("删除评论失败,remarkID: %v,err:%v", remarkID, err.Error())
		return err
	}
	global.Logger.Debug("删除评论: ", remarkID)
	return nil
}
