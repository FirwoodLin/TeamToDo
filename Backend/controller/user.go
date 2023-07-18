package controller

import (
	"TeamToDo/database"
	"TeamToDo/global"
	"TeamToDo/model/request"
	"TeamToDo/model/response"
	"TeamToDo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserRegistrationHandler(c *gin.Context) {
	var u request.UserRegisterRequest
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := validator.New().Struct(u); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	resp, err := database.UserRegister(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	if err := utils.PostEmail(u.Email, database.NewVerifyLinkUuid(u.Email)); err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("无法发送邮件"))
		global.Logger.Infof("无法发送邮件: %v", err)
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(*resp))
}

func UserVerifyHandler(c *gin.Context) {
	uuid := c.Query("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := database.VerifyEmail(uuid); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("激活账号失败"))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}

func UserLoginHandler(c *gin.Context) {
	var u request.UserSignInRequest
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	resp, err := database.UserSignIn(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	token := utils.GenerateJWTToken(strconv.FormatUint(uint64(resp.UserID), 10))
	if token == "" {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("无法颁发令牌"))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(map[string]interface{}{"user": *resp, "token": token}))
}

func UserUpdateInfoHandler(c *gin.Context) {
	var u request.UserUpdateRequest
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
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

func GetUserInfoHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	if userID == 0 {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("身份信息无效"))
		return
	}
	u, err := database.UserQueryOne(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(*u))
}
