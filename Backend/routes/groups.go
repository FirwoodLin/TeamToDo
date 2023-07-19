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

func registerAdvancedGroupRoutes(r *gin.RouterGroup) {
	advanced := r.Group("/:groupID", utils.MiddlewareRole(model.RoleAdmin))

	// 获取所有的加群申请
	advanced.GET("/applys", controller.GetApplysHandler)

	// 更新申请状态
	advanced.PUT("/applys/:applyID", controller.UpdateApplyStatusHandler)

	// // 更新成员状态
	// 更新成员状态这个太笼统了，设置管理员需要群主权限，移出成员只需要管理员权限
	// advanced.PUT
}
