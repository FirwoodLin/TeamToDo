package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"errors"
)

// ApplyGroupByID 申请加群(生成加群申请记录)
func ApplyGroupByID(userID uint, groupID uint) (*model.GroupApply, error) {
	var apply model.GroupApply
	apply.UserID = userID
	apply.GroupID = groupID
	apply.Status = model.ApplyStatusPending
	if err := global.Sql.Create(&apply).Error; err != nil {
		global.Logger.Errorf("数据库错误，申请加群失败，错误信息为：%v", err)
		return nil, err
	}
	return &apply, nil
}

// UpdateApplyStatus 更新申请状态（原同意加群）-仅限群主或者管理员- 同意申请；拒绝申请
func UpdateApplyStatus(applyID uint, applyStatus model.ApplyStatus) error {
	var apply model.GroupApply
	apply.GroupApplyID = applyID
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
	global.Logger.Debug("UpdateApplyStatus: apply: ", apply)
	return nil
}

// GetAllApplys 获取所有申请
func GetAllApplys(groupID uint) (applys []model.GroupApply, err error) {
	db := global.Sql.Model(&model.GroupApply{})
	if err = db.
		Where("groupID = ?", groupID).
		Find(&applys).
		Error; err != nil {
		global.Logger.Errorf("数据库错误，获取所有申请失败，错误信息为：%v", err)
		return nil, err
	}
	global.Logger.Debug("GetAllApplys: applys: ", applys)
	return applys, nil
}

// QueryApplyInfoByID 查询申请信息
func QueryApplyInfoByID(applyID uint) (apply model.GroupApply, err error) {
	db := global.Sql.Model(&model.GroupApply{})
	if err = db.Where("groupApplyID = ?", applyID).Find(&apply).Error; err != nil {
		global.Logger.Errorf("数据库错误，获取申请信息失败，错误信息为：%v", err)
		return apply, err
	}
	global.Logger.Debug("QueryApplyInfo: apply: ", apply)
	return apply, nil
}

//// FindGroupByUuid （加群-邀请链接）根据群UUID查找群ID
//func FindGroupByUuid(uuid string) (groupID uint, err error) {
//	// TODO
//	return 0, nil
//}
