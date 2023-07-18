package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"errors"
)

// UpdateApplyStatus 更新申请状态（原同意加群）
func UpdateApplyStatus(applyID uint, applyStatus model.ApplyStatus) error {
	var apply model.GroupApplication
	apply.GroupApplicationID = applyID
	if err := global.Sql.Find(&apply); err != nil {
		global.Logger.Errorf("数据库错误，查询申请(id:%v)失败，错误信息为：%v", applyID, err)
		return errors.New("该申请不存在")
	}
	if apply.Status != model.ApplyStatusPending {
		global.Logger.Infof("该申请已处理，申请ID为：%v", applyID)
		return errors.New("该申请已处理")
	}
	apply.Status = applyStatus
	if err := global.Sql.Save(&apply); err != nil {
		global.Logger.Errorf("数据库错误，更新申请状态(id:%v)失败，错误信息为：%v", applyID, err)
		return errors.New("数据库错误，更新申请状态失败")
	}
	return nil
}
