package controller

import (
	"TeamToDo/database"
	"TeamToDo/model"
	"TeamToDo/model/request"
	"TeamToDo/model/response"
	"TeamToDo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET: "/api/groups/:groupsID/applys"
// 获取所有的申请
func GetApplysHandler(c *gin.Context) {
	groupID := c.GetUint("groupID")

	// 获取所有的申请
	resp, err := database.GetAllApplys(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("群组不存在"))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(resp))
}

// PUT: "/api/groups/:groupID/applys/:applyID"
// 更新申请状态（同意，拒绝）
func UpdateApplyStatusHandler(c *gin.Context) {
	// groupID := c.GetUint("groupID")

	applyID, err := strconv.ParseUint(c.Param("applyID"), 10, 32)
	status, err2 := strconv.ParseInt(c.Query("status"), 10, 64)
	if err != nil || err2 != nil || model.ApplyStatus(status) < model.ApplyStatusPending || model.ApplyStatus(status) > model.ApplyStatusRejected {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	// 更新申请状态
	if err := database.UpdateApplyStatus(uint(applyID), model.ApplyStatus(status)); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	// 检查申请的状态，如果同意了，需要将用户加入群组
	if model.ApplyStatus(status) == model.ApplyStatusAgreed {
		// 查找申请ID对应的群组和用户ID
		apply, err := database.QueryApplyInfoByID(uint(applyID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("此申请已经不存在 "+err.Error()))
			return
		}
		// 将用户以普通身份加入群组
		if err := database.AddGroupMember(apply.UserID, apply.GroupID, model.RoleMember); err != nil {
			c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("此用户无法加入此群组"))
			return
		}
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}

// POST: "/api/groups/:groupID/join/codes"
// 创建并返回邀请码
func GetInviteCodeHandler(c *gin.Context) {
	groupID := c.GetUint("groupID")
	userID := c.GetUint("userID")

	var cr request.GroupCodeRequest
	if err := c.ShouldBind(&cr); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	cr.GroupID = groupID
	cr.InviterID = userID
	cr.Code = utils.GenerateInviteCode(request.InviteCode)
	cr.Type = request.InviteCode
	if err := database.InsertGroupJoinCode(&cr); err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("生成邀请码失败 "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(cr.Code))
}

// // POST: "/api/groups/:groupID/join/links"
// // 创建并返回邀请链接
// func GetInviteLinkHandler(c *gin.Context) {
// 	groupID := c.GetUint("groupID")
// 	userID := c.GetUint("userID")

// 	var cr request.GroupCodeRequest
// 	if err := c.ShouldBind(&cr); err != nil {
// 		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
// 		return
// 	}
// 	cr.GroupID = groupID
// 	cr.InviterID = userID
// 	// 前半代码完全相同
// 	cr.Code = utils.GenerateInviteCode(request.InviteLink)
// 	cr.Type = request.InviteLink
// 	if err := database.InsertGroupJoinCode(&cr); err != nil {
// 		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("生成邀请链接失败 "+err.Error()))
// 		return
// 	}
// 	// 这里要生成网址，注意从config里面读
// 	// 注意注意！！！这里没有完成！
// 	c.JSON(http.StatusOK, response.MakeSucceedResponse(cr.Code))
// }

// DELETE: "/api/groups/:groupID/members/:targetUserID"
// 移除群成员
func RemoveUserFromHandler(c *gin.Context) {
	groupID := c.GetUint("groupID")
	targetUserID, err := strconv.ParseUint(c.Param("targetUserID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := database.DeleteGroupMember(uint(targetUserID), groupID); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("无法移出群组 "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}
