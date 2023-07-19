package controller

import (
	"TeamToDo/database"
	"TeamToDo/model"
	"TeamToDo/model/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET: "/applys"
// 获取所有的申请
func GetApplysHandler(c *gin.Context) {
	// groupID := c.GetUint("groupID")

	// resp, err := database.
	// 获取所有申请好像还没写
}

// PUT: "/applys/:applyID"
// 更新申请状态（同意，拒绝）
func UpdateApplyStatusHandler(c *gin.Context) {
	// groupID := c.GetUint("groupID")

	applyID, err := strconv.ParseUint(c.Param("applyID"), 10, 32)
	status, err2 := strconv.ParseInt(c.Query("status"), 10, 64)
	if err != nil || err2 != nil || model.ApplyStatus(status) < model.ApplyStatusPending || model.ApplyStatus(status) > model.ApplyStatusRejected {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := database.UpdateApplyStatus(uint(applyID), model.ApplyStatus(status)); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	// 如果同意申请了，就将用户加入群组
	// 淦，又没有接口
	// 怎么感觉 给了很多接口，又什么也没有给啊
}
