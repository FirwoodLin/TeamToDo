package routes

import (
	"TeamToDo/controller"
	"TeamToDo/utils"

	"github.com/gin-gonic/gin"
)

func registerGroupsRoutes(e *gin.Engine) {
	group := e.Group("/api/groups", utils.MiddlewareJWTAuthorize())

	// 基础功能
	registerBasicGroupRoutes(group)

	// 高级功能
	// registerAdvancedGroupRoutes(group)

}

func registerBasicGroupRoutes(r *gin.RouterGroup) {
	// 通过群组ID加群
	r.POST("join", controller.JoinFromIDHandler)

	// 创建群组
	r.POST("/", controller.CreateGroupHandler)

	// 退出群组
	r.DELETE("/:groupID/members", controller.QuitGroupHandler)

	// 查看个人群组
	r.GET("/", controller.GetGroupsHandler)

}

// func registerAdvancedGroupRoutes(r *gin.RouterGroup) {
// 	advanced := r.Group("/:groupID", utils.MiddlewareRole(model.RoleAdmin))

// 	// // 获取所有的加群申请
// 	// advanced.GET("/applys", GetApplysHandler)

// 	// // 更新申请状态
// 	// advanced.PUT("/applys/:applyID", UpdateApplyStatusHandler)

// 	//
// }
