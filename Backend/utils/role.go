package utils

import (
	"TeamToDo/global"
	"TeamToDo/model"
	"TeamToDo/model/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CheckUserInGroup 检查用户是否在群组中！！！（重要）（附带检查用户在群组中的身份）
func CheckUserInGroup(userID uint, groupID uint) model.Role {
	var userGroup model.UserGroup
	_ = global.Sql.Where("user_id = ? AND group_id = ?", userID, groupID).First(&userGroup)
	return userGroup.Role
}

func MiddlewareRole(roleRequested model.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("userID")
		groupID, err := strconv.ParseUint(c.Param("groupID"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.InvalidInfoError)
			c.Abort()
			return
		}
		role := CheckUserInGroup(userID, uint(groupID))
		if role < roleRequested {
			c.JSON(http.StatusUnauthorized, response.MakeFailedResponse("权限不足"))
			c.Abort()
			return
		}
		c.Set("groupID", uint(groupID))
		c.Set("role", role)
	}
}
