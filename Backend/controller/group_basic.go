package controller

import (
	"TeamToDo/database"
	"TeamToDo/global"
	"TeamToDo/model"
	"TeamToDo/model/request"
	"TeamToDo/model/response"
	"TeamToDo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// POST: "/api/groups/join"
// 通过群ID加群
func JoinFromIDHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	groupID, err := strconv.ParseUint(c.PostForm("groupID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 创建申请时大概会检查该群是否存在
	resp, err := database.ApplyGroupByID(userID, uint(groupID))

	/***************************/
	// 应要求，现在申请通过群ID加群会直接加入
	database.UpdateApplyStatus(resp.GroupApplyID, model.ApplyStatusAgreed)
	database.AddGroupMember(resp.UserID, resp.GroupID, model.RoleMember)
	/***************************/

	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(gin.H{"apply": *resp, "status": "加入成功"}))
}

// POST: "/api/groups/join/codes"
// 通过群组邀请码加群
func JoinFromCodeHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	code := c.PostForm("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	resp, err := database.QueryGroupJoinInfo(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("不存在的邀请码"))
		return
	}
	if err := database.AddGroupMember(userID, resp.GroupID, model.RoleMember); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("该用户无法加入群组 "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}

// GET: "/api/groups/join/links"
// 通过群组邀请链接加群
// 暂时留空

// POST: "/api/groups"
// 创建群聊
func CreateGroupHandler(c *gin.Context) {
	userID := c.GetUint("userID")

	var g request.GroupCreateRequest

	if err := c.ShouldBind(&g); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 创建新的群
	resp, err := database.GroupCreate(&g)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	// 群主一定要进入群聊，否则是无效群
	if err := database.AddGroupMember(userID, resp.GroupID, model.RoleOwner); err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("创建群聊失败"))
		global.Logger.Infof("群主没有正确地加入群聊")
		// 这里理论上应该删除这个群，否则会在数据库中留有无效数据，但是没给接口
		return
	}

	c.JSON(http.StatusOK, response.MakeSucceedResponse(*resp))
}

// DELETE: "api/groups/:groupID/members"
// 退出群聊
func QuitGroupHandler(c *gin.Context) {
	userID := c.GetUint("userID")

	groupID, err := strconv.ParseUint(c.Param("groupID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 检查用户在群组中的权限
	role := utils.CheckUserInGroup(userID, uint(groupID))

	if role == model.RoleVisitor {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("你并不在群聊中"))
		return
	} else if role == model.RoleOwner {
		DisbandGroup(c, uint(groupID))
		return
	}
	if err := database.QuitGroup(userID, uint(groupID)); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("无法退出这个群"))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}

// GET: "/api/groups"
// 获得当前用户的所有群聊
func GetGroupsHandler(c *gin.Context) {
	userID := c.GetUint("userID")

	// 直接查询该用户所在的所有群聊
	resp, err := database.FindUserJoinedGroups(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	groups := make([]model.Group, 0)
	for _, ug := range *resp {
		groups = append(groups, ug.Group)
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(gin.H{"groups": groups}))
}

// GET: "/api/groups/:groupID/members"
// 查看群组所有成员
func GetAllUsersInGroupHandler(c *gin.Context) {
	groupID, err := strconv.ParseUint(c.Param("groupID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 查询这个群组的所有成员
	resp, err := database.FindGroupMembers(uint(groupID))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("群组不存在"))
		return
	}
	users := make([]response.UserGroupResponse, 0)
	for _, ug := range resp {
		var ur response.UserGroupResponse
		if err := copier.Copy(&ur, &ug); err != nil {
			c.JSON(http.StatusBadRequest, response.MakeFailedResponse("结构拷贝错误"))
			return
		}
		users = append(users, ur)
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(gin.H{"members": users}))
}

// GET: "/api/groups/:groupID/role"
// 查询自己在当前群组中的角色
func GetSelfRoleInGroupHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	groupID, err := strconv.ParseUint(c.Param("groupID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(gin.H{"role": utils.CheckUserInGroup(userID, uint(groupID))}))
}
