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
)

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
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(*resp))
}

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

// 退出群聊
func QuitGroupHandler(c *gin.Context) {
	userID := c.GetUint("userID")

	groupID, err := strconv.ParseUint(c.Param("groupID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 退出群聊的前提是用户在群聊中
	// 这里没有检查用户是不是群主，如果用户是群主会涉及群主转让，直接删除群聊等。。。
	// 之后一定要加上检查是不是群主的函数
	if utils.CheckUserInGroup(userID, uint(groupID)) == model.RoleVisitor {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("你并不在群聊中"))
		return
	}
	if err := database.QuitGroup(userID, uint(groupID)); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}

// 获得当前用户的所有群聊
func GetGroupsHandler(c *gin.Context) {
	userID := c.GetUint("userID")

	// 直接查询该用户所在的所有群聊
	resp, err := database.FindUserJoinedGroups(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(*resp))
}
