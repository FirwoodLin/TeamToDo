package controller

import (
	"TeamToDo/database"
	"TeamToDo/model"
	"TeamToDo/model/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PUT: "/api/groups/:groupID/members/:targetUserID"
// 设置/取消管理员权限
func UpdateUserRoleHandler(c *gin.Context) {
	groupID := c.GetUint("groupID")
	role := model.Role(c.GetInt("role"))
	if role != model.RoleOwner {
		c.JSON(http.StatusUnauthorized, response.MakeFailedResponse("需要群主权限"))
		c.Abort()
		return
	}
	targetUserID, err := strconv.ParseUint(c.Param("targetUserID"), 10, 32)
	targetUserRole, err2 := strconv.ParseInt(c.PostForm("role"), 10, 64)
	if err != nil || err2 != nil || (targetUserRole != int64(model.RoleMember) && targetUserRole != int64(model.RoleAdmin)) {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := database.UpdateMemberRole(uint(targetUserID), groupID, model.Role(targetUserRole)); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("修改权限失败"))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}
