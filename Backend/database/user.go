package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"TeamToDo/model/request"
	"TeamToDo/model/response"
	"TeamToDo/utils"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

// UserRegister 用户注册 - 调用/查重/函数和/创建/用户函数
func UserRegister(userReq *request.UserRegisterRequest) (userResponse *response.UserResponse, err error) {
	// 检验数据是否合规
	validate := validator.New()
	if err := validate.Struct(userReq); err != nil {
		global.Logger.Info("注册用户时,数据不合规\n")
		return nil, err
	}
	// 创建用户
	userResponse, err = createUser(userReq)
	if err != nil {
		global.Logger.Errorf("创建用户错误\n")
		return nil, err
	}
	return userResponse, nil
}

// createUser 创建用户 - 内部函数
func createUser(userReq *request.UserRegisterRequest) (*response.UserResponse, error) {
	// 改用 copier 进行结构体转换
	user := model.User{}
	err := copier.Copy(&user, userReq)
	if err != nil {
		global.Logger.Errorf("用户注册时,结构体转换错误\n")
		return nil, err
	}
	// 密码加密
	err = utils.EncryptUserPassword(&user.Password)
	if err != nil {
		global.Logger.Errorf("密码加密错误: %v", err)
		return nil, err
	}
	// 存入数据库
	if err := global.Sql.Create(&user).Error; err != nil {
		global.Logger.Errorf("创建用户错误\n")
		return nil, err
	}
	// 赋予用户名和头像
	width := 8
	user.UserName = fmt.Sprintf("User%0*s", width, strconv.FormatInt(int64(user.UserID), 10))
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndNum := rng.Intn(30)
	user.UserAvatar = fmt.Sprintf(global.Server.Avatar.UserUrl, rndNum)
	// 进行更新
	if err := global.Sql.Model(&user).Updates(user).Error; err != nil {
		global.Logger.Errorf("更新用户名和头像错误\n")
		return nil, err
	}
	// 进行返回
	var userResponse response.UserResponse
	err = copier.Copy(&userResponse, &user)
	if err != nil {
		global.Logger.Errorf("用户注册时,结构体转换错误\n")
		return nil, err
	}
	return &userResponse, nil
}

// UserSignIn 用户登陆； 1.检查邮箱是否存在 2.邮箱密码是否一致
func UserSignIn(userReq *request.UserSignInRequest) (userResponse *response.UserResponse, err error) {
	// 只有登陆的时候可以产生 JWT token
	// 检查邮箱是否存在
	var user model.User
	result := global.Sql.Where("email = ?", userReq.Email).First(&user)
	if result.RowsAffected == 0 {
		global.Logger.Infof("登陆邮箱不存在%v\n", userReq.Email)
		return nil, errors.New("该邮箱未注册")
	}
	// 检查密码是否一致
	if err := utils.ComparePassword(userReq, &user); err != nil {
		//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password)); err != nil {
		global.Logger.Infof("登陆密码不正确,email:%v,pwd:%v\n", userReq.Email, userReq.Password)
		return nil, errors.New("密码不正确")
	}
	userResponse = &response.UserResponse{UserID: user.UserID}
	global.Logger.Infof("登陆成功,id：%v\n", userResponse.UserID)
	// 检查是否激活
	if !user.IsVerified {
		global.Logger.Infof("用户未激活,id：%v\n", userResponse.UserID)
		return nil, errors.New("用户未激活")
	}
	_ = copier.Copy(&userResponse, &user)
	// 返回用户信息
	return userResponse, nil
}

// UserUpdate 用户更新资料；将 UserID 在 userReq 结构体中传递
func UserUpdate(userReq *request.UserUpdateRequest) error {
	user := &model.User{}
	user.UserID = userReq.UserID
	db := global.Sql.Model(&model.User{})
	// 查找用户
	err := db.Find(&user).Error
	if err != nil {
		global.Logger.Infof("查找不到用户,id:%v", userReq.UserID)
		return err
	}
	// 更新密码
	if userReq.NewPassword != "" {
		// 检查旧密码是否正确
		userSignInReq := &request.UserSignInRequest{Password: userReq.OldPassword}
		//TODO：重构Compare 函数，只传入密码，不传入结构体
		err := utils.ComparePassword(userSignInReq, user)
		if err != nil {
			global.Logger.Infof("旧密码不正确,id:%v", userReq.UserID)
			return err
		}
		// 加密新密码
		_ = utils.EncryptUserPassword(&userReq.NewPassword)
		user.Password = userReq.NewPassword
	}
	// 更新用户名
	user.UserName = userReq.UserName
	// 存入数据库
	//global.Sql.Model(user).Updates(userReq)
	global.Sql.Save(&user)
	return nil
}

// UserQueryOne 查询用户信息 -- 传入用户 ID;返回三要素
func UserQueryOne(userID uint) (*response.UserResponse, error) {
	var userResponse response.UserResponse
	var user model.User
	db := global.Sql.Model(&model.User{})
	// 使用手动 SELECT 方法在面临字段更改时会很麻烦
	user.UserID = userID
	if err := db.Find(&user).Error; err != nil {
		global.Logger.Infof("查询用户信息,查找不到用户,id:%v", userID)
		return nil, err
	}
	// 结构体转换，返回
	_ = copier.Copy(&userResponse, &user)
	return &userResponse, nil
}

// Deprecated: DeleteUser is deprecated. 暂时用不到
func DeleteUser(userID uint) (err error) {
	// DeleteUser 删除用户 = 注销账号会用到 ** 可能会用到
	err = global.Sql.Delete(&model.User{}, userID).Error
	if err != nil {
		global.Logger.Errorf("DeleteUser 出错,%v\n", err.Error())
		return err
	}
	//log.Printf("[info]model-DeleteUser,delete userid:%v\n", userID)

	return nil
}

// UserQueryOneAllInfo 查询用户信息 -- 传入用户 ID;返回所有信息
func UserQueryOneAllInfo(userID uint) (*model.User, error) {
	var user model.User
	db := global.Sql.Model(&model.User{})
	user.UserID = userID
	if err := db.Find(&user).Error; err != nil {
		global.Logger.Infof("查询用户信息,查找不到用户,id:%v", userID)
		return nil, err
	}
	return &user, nil
}

// QueryUserByEmail 根据邮箱查询用户信息
func QueryUserByEmail(email string) (*model.User, error) {
	var user model.User
	db := global.Sql.Model(&model.User{})
	if err := db.Where("email = ?", email).Find(&user).Error; err != nil {
		global.Logger.Infof("查询用户信息,查找不到用户,email:%v", email)
		return nil, err
	}
	return &user, nil
}
