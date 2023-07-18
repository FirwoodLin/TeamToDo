package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"TeamToDo/utils"
	"time"
)

// CheckIsVerified 检查用户是否已经验证过邮箱
// deprecated: 在登陆时完成
func CheckIsVerified() {

}

// VerifyEmail 接收 uuid,激活 uuid 对应的邮箱
func VerifyEmail(uuid string) error {
	db := global.Sql.Model(&model.EmailVerification{})
	var emailVerification model.EmailVerification
	// 查找 没有过期的 uuid;利用 CreatedAt省去 IsUsed 字段
	if err := db.Where("uuid = ?", uuid).
		Where("created_at > ?", time.Now().Add(-time.Hour)).
		Find(&emailVerification).
		Error; err != nil {
		global.Logger.Infof("uuid 不存在或已过期,uuid: %v,err:%v", uuid, err.Error)
		return err
	}
	// 删除原有的 uuid
	if err := db.Where("uuid = ?", uuid).Delete(&emailVerification).Error; err != nil {
		global.Logger.Errorf("删除 uuid 失败,uuid: %v,err:%v", uuid, err.Error())
	}
	db = global.Sql.Model(&model.User{})
	// 更新用户激活状态
	if err := db.
		Where("email = ?", emailVerification.Email).
		Update("is_verified", true).
		Error; err != nil {
		global.Logger.Infof("更新用户激活状态失败,uuid: %v,err:%v", uuid, err.Error())
		return err
	}
	return nil
}

func NewVerifyLinkUuid(email string) string {
	uuid := utils.NewUuid()
	emailVerification := model.EmailVerification{
		Uuid:  uuid,
		Email: email,
	}
	if err := global.Sql.Create(&emailVerification).Error; err != nil {
		global.Logger.Errorf("创建 uuid 失败,uuid: %v,err:%v", uuid, err.Error())
		return ""
	}
	return uuid
}
