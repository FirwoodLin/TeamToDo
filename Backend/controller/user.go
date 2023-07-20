package controller

import (
	"TeamToDo/database"
	"TeamToDo/model/request"
	"TeamToDo/model/response"
	"TeamToDo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 用户注册处理
func UserRegistrationHandler(c *gin.Context) {
	var u request.UserRegisterRequest
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 检查信息的合法性
	if err := validator.New().Struct(u); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 数据库插入后返回信息
	resp, err := database.UserRegister(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	// 发送邮件
	if err := utils.PostEmail(u.Email, utils.GenerateActivateMail(database.NewVerifyLinkUuid(u.Email))); err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("无法发送邮件"))
		return
	}
	// 注册成功，等待激活
	c.JSON(http.StatusOK, response.MakeSucceedResponse(*resp))
}

// 用户账号验证处理
func UserVerifyHandler(c *gin.Context) {
	// 接受邮件中传入的uuid
	uuid := c.Query("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 对数据库中的账号进行验证和激活
	if err := database.VerifyEmail(uuid); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("激活账号失败"))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}

// 用户登录处理
func UserLoginHandler(c *gin.Context) {
	var u request.UserSignInRequest
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 数据库查询用户信息
	resp, err := database.UserSignIn(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	// 登录成功，生成令牌，转换成string后传入
	token := utils.GenerateJWTToken(strconv.FormatUint(uint64(resp.UserID), 10))
	if token == "" {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("无法颁发令牌"))
		return
	}
	// 返回用户信息及其令牌
	c.JSON(http.StatusOK, response.MakeSucceedResponse(map[string]interface{}{"user": *resp, "token": token}))
}

// 更新用户信息
func UserUpdateInfoHandler(c *gin.Context) {
	var u request.UserUpdateRequest
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 维护userID
	u.UserID = c.GetUint("userID")
	// 全部传进去
	if err := database.UserUpdate(&u); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}

func UserUpdateAvatorHandler(c *gin.Context) {

}

func GetPublicKeyHandler(c *gin.Context) {

}

// 获取用户信息
func GetUserInfoHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	// 在数据库中查询用户信息
	u, err := database.UserQueryOne(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	// 返回数据库中查询的信息
	c.JSON(http.StatusOK, response.MakeSucceedResponse(*u))
}
