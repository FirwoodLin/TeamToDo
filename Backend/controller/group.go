package controller

import (
	"TeamToDo/database"
	"TeamToDo/model/request"
	"TeamToDo/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JoinFromIDHandler(c *gin.Context) {
	// userID := c.GetUint("userID")
	// groupID, err := strconv.ParseUint(c.PostForm("groupID"), 10, 32)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, response.InvalidInfoError)
	// 	return
	// }

}

func CreateGroupHandler(c *gin.Context) {
	// userID := c.GetUint("userID")

	var g request.GroupCreateRequest

	if err := c.ShouldBind(&g); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	resp, err := database.GroupCreate(&g)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	// 这里群主一定要在群里，等待明天的更新~
	c.JSON(http.StatusOK, response.MakeSucceedResponse(*resp))
}

func QuitGroupHandler(c *gin.Context) {
	// userID := c.GetUint("userID")

}

func GetGroupsHandler(c *gin.Context) {
	// userID := c.GetUint("userID")

}
