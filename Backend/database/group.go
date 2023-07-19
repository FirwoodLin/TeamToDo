package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"TeamToDo/model/request"

	"github.com/jinzhu/copier"
)

// GroupCreate 	根据请求创建群组
func GroupCreate(groupReq *request.GroupCreateRequest) (*model.Group, error) {
	var group model.Group
	err := copier.Copy(&group, &groupReq)
	if err != nil {
		global.Logger.Infof("创建 Group 时,结构体转换错误\n")
		return nil, err
	}
	group.GroupAvatar = global.Server.Avatar.GroupUrl
	err = global.Sql.Create(&group).Error
	if err != nil {
		global.Logger.Errorf("创建 Group 时,数据库错误\n")
		return nil, err
	}
	return &group, err
}


// CheckUserInGroup 检查用户是否在群组中（附带检查用户在群组中的身份）
func CheckUserInGroup(userID uint, groupID uint) model.Role {
	var userGroup model.UserGroup
	_ = global.Sql.Where("user_id = ? AND group_id = ?", userID, groupID).First(&userGroup)
	return userGroup.Role
}

// FindUserJoinedGroups 查看个人的群组；查找用户ID加入的所有群组
func FindUserJoinedGroups(userID uint) (*[]model.UserGroup, error) {
	db := global.Sql
	var userGroups []model.UserGroup
	if err := db.Where("userID = ?", userID).Preload("Group").Find(&userGroups).Error; err != nil {
		// 处理错误
		global.Logger.Infof("查找用户加入的所有群组时,数据库错误\n")
	}
	return &userGroups, nil
}

// FindGroupMembers 查找群组中的所有成员
func FindGroupMembers(groupID uint) (*[]model.UserGroup, error) {
	db := global.Sql
	var userGroups []model.UserGroup
	if err := db.Where("groupID = ?", groupID).Preload("User").Find(&userGroups).Error; err != nil {
		// 处理错误
		global.Logger.Infof("查找群组中的所有成员时,数据库错误\n")
	}
	return &userGroups, nil
}

// AddGroupMember 添加群组成员 - 服务层调用
func AddGroupMember(userID uint, groupID uint) error {
	db := global.Sql
	var userGroup model.UserGroup
	userGroup.UserID = userID
	userGroup.GroupID = groupID
	//userGroup.Role = role
	if err := db.Create(&userGroup).Error; err != nil {
		// 处理错误
		global.Logger.Infof("添加群组成员时,数据库错误\n")
		return err
	}
	return nil
}

// QuitGroup 退出群组
func QuitGroup(userID uint, groupID uint) error {
	db := global.Sql
	var userGroup model.UserGroup
	if err := db.Where("userID = ? AND groupID = ?", userID, groupID).First(&userGroup).Error; err != nil {
		// 处理错误
		global.Logger.Infof("退出群组时,数据库错误，查询不到记录\n")
		return err
	}
	if err := db.Delete(&userGroup).Error; err != nil {
		// 处理错误
		global.Logger.Infof("退出群组时,数据库错误\n")
		return err
	}
	return nil
}
