package database

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"TeamToDo/model/request"
	"TeamToDo/model/response"
	"TeamToDo/utils"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// RegisterUser 注册用户 - 调用*查重*函数和*创建*用户函数
func RegisterUser(userReq *request.UserRegisterRequest) (uid uint, err error) {
	// 检验数据是否合规
	validate := validator.New()
	if err := validate.Struct(userReq); err != nil {
		global.Logger.Info("注册用户时,数据不合规\n")
		return 0, err
	}
	// 查询该用户是否注册过(名称/邮箱)
	if isExistUser(userReq.UserName, userReq.Email) {
		return 0, errors.New("该用户名/邮箱已注册")
	}
	// 注册用户
	uid, err = createUser(userReq)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

// isExistUser 检查用户名和邮箱是否注册过
func isExistUser(name, email string) bool {
	var user model.User
	result := global.Sql.Where("name = ? OR email = ?", name, email).First(&user)
	// 如果没有找到记录，返回 false
	// 静态检查建议：直接返回布尔值，不进行 if 判断
	return result.RowsAffected != 0
}

// createUser 创建用户
func createUser(userReq *request.UserRegisterRequest) (uint, error) {
	// 改用 copier 进行结构体转换
	user := &model.User{}
	err := copier.Copy(&user, &userReq)
	if err != nil {
		global.Logger.Errorf("用户注册时,结构体转换错误\n")
		return 0, err
	}
	// 密码加密
	err = utils.EncryptUserPassword(&user.Password)
	if err != nil {
		global.Logger.Errorf("密码加密错误: %v", err)
		return 0, err
	}
	if err := global.Sql.Create(user).Error; err != nil {
		global.Logger.Errorf("创建用户错误\n")
		return 0, err
	}
	return user.UserID, nil
}

// CheckUser 用户登录时 检查邮箱是否存在 邮箱密码是否一致 **是否是管理员**
func CheckUser(userSReq *request.UserSignInRequest) (userResponse *response.UserSignInResponse, err error) {
	// 检查邮箱是否存在
	var user model.User
	result := global.Sql.Where("email = ?", userSReq.Email).First(&user)
	if result.RowsAffected == 0 {
		global.Logger.Infof("登陆邮箱不存在%v\n", userSReq.Email)
		return nil, errors.New("该邮箱未注册")
	}
	// 检查密码是否一致
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userSReq.Password)); err != nil {
		global.Logger.Infof("登陆密码不正确%v\n", userSReq.Email)
		return nil, errors.New("密码不正确")
	}
	userResponse = &response.UserSignInResponse{UserID: user.UserID}
	global.Logger.Infof("登陆成功,id：%v\n", userResponse.UserID)

	return userResponse, nil
}

// UpdateUser 将 UserID 在 userReq 结构体中传递
func UpdateUser(userReq *request.UserUpdateRequest) {
	user := &model.User{}
	user.UserID = userReq.UserID
	if len(userReq.Password) != 0 {
		_ = utils.EncryptUserPassword(&userReq.Password)
	}
	global.Sql.Model(user).Updates(userReq)
}

// QueryOneUser 查询某个用户的信息 -- 用户头像和ID
func QueryOneUser(userID uint) (*response.UserQueryResponse, error) {
	var userResponse response.UserQueryResponse
	var user model.User
	db := global.Sql.Model(&model.User{})
	// 使用手动 SELECT 方法在面临字段更改时会很麻烦
	user.UserID = userID
	if err := db.Find(&user).Error; err != nil {
		global.Logger.Infof("查找不到用户,id:%v", userID)
		return nil, err
	}
	_ = copier.Copy(&userResponse, &user)
	return &userResponse, nil
}

// DeleteUser 删除用户 = 注销账号会用到 ** 可能会用到
func DeleteUser(userID uint) (err error) {
	err = global.Sql.Delete(&model.User{}, userID).Error
	if err != nil {
		global.Logger.Errorf("DeleteUser 出错,%v\n", err.Error())
		return err
	}
	log.Printf("[info]model-DeleteUser,delete userid:%v\n", userID)

	return nil
}
