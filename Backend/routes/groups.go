package routes

import (
	"TeamToDo/controller"
	"TeamToDo/model"
	"TeamToDo/utils"

	"github.com/gin-gonic/gin"
)

func registerGroupsRoutes(e *gin.Engine) {
	group := e.Group("/api/groups", utils.MiddlewareJWTAuthorize())

	// 基础功能
	registerBasicGroupRoutes(group)

	// 高级功能
	registerAdvancedGroupRoutes(group)

}

func registerBasicGroupRoutes(r *gin.RouterGroup) {
	// 通过群组ID加群
	r.POST("/join", controller.JoinFromIDHandler)

	// 通过群组邀请码加群
	r.POST("/join/codes", controller.JoinFromCodeHandler)

	// // 通过群组邀请链接加群
	// // 暂时留空
	// r.POST("/join/links", controller.JoinFromLinkHandler)

	// 创建群组
	r.POST("/", controller.CreateGroupHandler)

	// 退出群组
	r.DELETE("/:groupID/members", controller.QuitGroupHandler)

	// 查看个人群组
	r.GET("/", controller.GetGroupsHandler)

	// 查询群组所有成员
	r.GET("/:groupID/members", controller.GetAllUsersInGroupHandler)

	// 查询自己在当前群组的权限
	r.GET("/:groupID/role", controller.GetSelfRoleInGroupHandler)
}

func registerAdvancedGroupRoutes(r *gin.RouterGroup) {
	advanced := r.Group("/:groupID", utils.MiddlewareRole(model.RoleAdmin))

	// 获取所有的加群申请
	advanced.GET("/applys", controller.GetApplysHandler)

	// 更新申请状态
	advanced.PUT("/applys/:applyID", controller.UpdateApplyStatusHandler)

	// 更新成员状态（设置/取消管理员）
	// 群主功能，考虑是不是顶级功能
	advanced.PUT("/members/:targetUserID", controller.UpdateUserRoleHandler)

	// 移出成员
	advanced.DELETE("/member/:targetUserID", controller.RemoveUserFromHandler)

	// 创建邀请码
	advanced.POST("/join/codes", controller.GetInviteCodeHandler)

	// // 创建邀请链接
	// // 链接部分暂时留空
	// advanced.POST("/join/links", controller.GetInviteLinkHandler)
}
