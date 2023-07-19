package routes

import (
	"TeamToDo/controller"
	"TeamToDo/utils"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(e *gin.Engine) {
	user := e.Group("/api/users")

	// 用户注册
	user.POST("/registration", controller.UserRegistrationHandler)

	// 用户激活
	user.GET("/verify", controller.UserVerifyHandler)

	// 用户登录
	user.POST("/login", controller.UserLoginHandler)

	user.Use(utils.MiddlewareJWTAuthorize())
	{
		// 用户更新信息
		user.PUT("", controller.UserUpdateInfoHandler)

		// 用户更新头像
		user.POST("/avator", controller.UserUpdateAvatorHandler)

		// 用户获取公钥
		user.GET("/publicKey", controller.GetPublicKeyHandler)

		// 查询用户信息
		user.GET("/:userID", controller.GetUserInfoHandler)
	}
}
