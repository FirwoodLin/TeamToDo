package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"TeamToDo/model/request"
	"gorm.io/gorm"

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
// 使用中 暂时不修改参数类型
func FindUserJoinedGroups(userID uint) (*[]model.UserGroup, error) {
	db := global.Sql
	var userGroups []model.UserGroup
	if err := db.
		Where("userID = ?", userID).
		Preload("Group").
		Preload("User").
		Find(&userGroups).Error; err != nil {
		// 处理错误
		global.Logger.Infof("查找用户加入的所有群组时,数据库错误\n")
	}
	return &userGroups, nil
}

// FindGroupMembers 查找群组中的所有成员
func FindGroupMembers(groupID uint) (userGroups []model.UserGroup, err error) {
	db := global.Sql
	if err := db.Where("groupID = ?", groupID).Preload("User").Find(&userGroups).Error; err != nil {
		// 处理错误
		global.Logger.Infof("查找群组中的所有成员时,数据库错误\n")
		return nil, err
	}
	return userGroups, nil
}

// AddGroupMember 添加群组成员 - 服务层调用
func AddGroupMember(userID uint, groupID uint, role model.Role) error {
	userGroup := model.UserGroup{
		UserID:  userID,
		GroupID: groupID,
		Role:    role,
	}
	tx := global.Sql.Begin() // 由于涉及两个表的写操作，使用事务
	if err := tx.Create(&userGroup).Error; err != nil {
		// 处理错误
		global.Logger.Infof("添加群组成员,新建UserGroup记录时,数据库错误\n")
		tx.Rollback()
		return err
	}
	if err := tx.Model(&model.Group{}).
		Where("groupID = ?", groupID).
		Update("memberCount", gorm.Expr("memberCount + ?", 1)).
		Error; err != nil {
		// 处理错误
		global.Logger.Infof("添加群组成员，更新 memberCount 时数据库错误\n")
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// UpdateMemberRole 更新成员角色（不包括踢出）
func UpdateMemberRole(userID uint, groupID uint, role model.Role) error {
	db := global.Sql
	// 查找 userGroup 记录
	var userGroup model.UserGroup
	if err := db.Where("userID = ? AND groupID = ?", userID, groupID).
		First(&userGroup).Error; err != nil {
		// 处理错误
		global.Logger.Infof("更新成员状态时,数据库错误，查询不到记录\n")
		return err
	}
	userGroup.Role = role
	// 更新角色
	if err := db.Save(&userGroup).Error; err != nil {
		// 处理错误
		global.Logger.Infof("更新成员状态时,数据库错误\n")
		return err
	}
	return nil
}

// DeleteGroupMember 删除群组成员 - 踢出成员/退出群组通用
func DeleteGroupMember(userID uint, groupID uint) error {
	tx := global.Sql.Begin()
	// 查找记录
	var userGroup model.UserGroup
	if err := tx.Where("userID = ? AND groupID = ?", userID, groupID).
		First(&userGroup).Error; err != nil {
		// 处理错误
		global.Logger.Infof("删除群组成员时,数据库错误，查询不到记录\n")
		tx.Rollback()
		return err
	}
	// 删除记录
	if err := tx.Delete(&userGroup).Error; err != nil {
		// 处理错误
		global.Logger.Infof("删除群组成员时,数据库错误\n")
		tx.Rollback()
		return err
	}
	// 更新群组成员数量
	if err := tx.Model(&model.Group{}).
		Where("group_id = ?", groupID).
		Update("memberCount", gorm.Expr("memberCount - ?", 1)).
		Error; err != nil {
		// 处理错误
		global.Logger.Infof("删除群组成员时,更新 memberCount 时数据库错误\n")
		tx.Rollback()
		return err
	}
	// 完成事务，提交
	tx.Commit()
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

// DeleteUserGroups 根据群组ID,删除所有UserGroup关系
func DeleteUserGroups(groupID uint) error {
	db := global.Sql
	if err := db.Where("groupID = ?", groupID).
		Delete(&model.UserGroup{}).Error; err != nil {
		// 处理错误
		global.Logger.Infof("删除群组时,数据库错误\n")
		return err
	}
	return nil
}

// DeleteGroup 删除群组
func DeleteGroup(groupID uint) error {
	db := global.Sql
	if err := db.Where("groupID = ?", groupID).
		Delete(&model.Group{}).Error; err != nil {
		// 处理错误
		global.Logger.Infof("删除群组时,数据库错误\n")
		return err
	}
	return nil
}
