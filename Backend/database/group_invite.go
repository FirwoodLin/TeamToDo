package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"TeamToDo/model/request"
	"github.com/jinzhu/copier"
)

// InsertGroupJoinCode 插入群邀请码(创建邀请码API使用);预先填好邀请码；群ID；邀请人ID
func InsertGroupJoinCode(groupReq *request.GroupCodeRequest) (err error) {
	var groupJoinCode model.GroupJoinCode
	if err := copier.Copy(&groupJoinCode, &groupReq); err != nil {
		global.Logger.Errorf("创建 GroupJoinCode 时,结构体转换错误\n")
		return err
	}
	if err := global.Sql.Save(&groupJoinCode).Error; err != nil {
		global.Logger.Errorf("邀请码插入错误,%v\n", err)
	}
	return nil
}

// DeleteGroupJoinCode 删除群邀请码(创建邀请码API使用);预先填好邀请码；群ID；邀请人ID
func DeleteGroupJoinCode(code string) (err error) {
	// 删除之前检验有没有权限删除（加群-邀请码 相关）
	err = global.Sql.
		Where("code = ?", code).
		Delete(&model.GroupJoinCode{}).Error
	if err != nil {
		global.Logger.Infof("删除群邀请码失败，错误信息为：%v", err)
		return err
	}
	return nil
}

// QueryGroupJoinInfo （加群-邀请码）根据*群邀请码*查找邀码相关信息（群ID,邀请人ID）
func QueryGroupJoinInfo(code string) (*model.GroupJoinCode, error) {
	// code: 群邀请码
	var groupJoinCode model.GroupJoinCode
	if err := global.Sql.Where("code = ?", code).First(&groupJoinCode).Error; err != nil {
		global.Logger.Infof("查询群邀请码失败，错误信息为：%v", err)
		return nil, err
	}
	return &groupJoinCode, nil
}

// GetAllGroupCode 查询群邀请码；
func GetAllGroupCode(groupID uint, inviterID uint) (*[]model.GroupJoinCode, error) {
	// @param groupID 群ID
	// @param inviterID 邀请人ID;如果为 0,代表群主查看所有邀请码；如果不为 0,代表管理员查看自己的邀请码
	var groupJoinCodes []model.GroupJoinCode
	if inviterID == 0 {
		// 群主查看所有邀请码
		if err := global.Sql.Where("group_id = ?", groupID).Find(&groupJoinCodes).Error; err != nil {
			global.Logger.Infof("查询群邀请码失败，错误信息为：%v", err)
			return nil, err
		}
	} else {
		// 管理员查看自己的邀请码
		if err := global.Sql.Where("group_id = ? AND inviter_id = ?", groupID, inviterID).Find(&groupJoinCodes).Error; err != nil {
			global.Logger.Infof("数据库错误，查询群邀请码失败，错误信息为：%v", err)
			return nil, err
		}
	}
	return &groupJoinCodes, nil
}
