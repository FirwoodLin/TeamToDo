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
	"github.com/jinzhu/copier"
)

// POST: "/api/tasks"
// 新建任务
func CreateTaskHandler(c *gin.Context) {
	var (
		tr   request.TaskRequest
		t    model.Task
		resp response.TaskResponse
	)
	// 解析参数
	userID := c.GetUint("userID")
	if err := c.ShouldBind(&tr); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	tr.OwnerID = userID
	// 鉴权
	if utils.CheckUserInGroup(userID, tr.GroupID) == model.RoleVisitor {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("你并不在该群组中"))
		return
	}
	// 结构转换
	if err := copier.Copy(&t, &tr); err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("结构转换错误"))
		return
	}
	// 创建任务
	if err := database.TaskCreate(&t); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("创建任务失败 "+err.Error()))
		return
	}
	if err := copier.Copy(&resp, &t); err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("结构转换错误"))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(gin.H{"task": resp}))
}

// DELETE: "/api/tasks/:taskID"
// 删除任务
func RemoveTaskHandler(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("taskID"), 10, 32)
	userID := c.GetUint("userID")

	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	t, err := database.QueryTaskByTaskID(uint(taskID))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("找不到任务"))
		return
	}
	if t.OwnerID != userID && utils.CheckUserInGroup(userID, t.GroupID) <= model.RoleMember {
		c.JSON(http.StatusUnauthorized, response.MakeFailedResponse("没有删除这个任务的权限"))
		return
	}
	if err := database.DeleteTask(uint(taskID)); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("任务删除失败 "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
}

// PUT: "/api/tasks/:taskID"
// 修改任务
func UpdateTaskHandler(c *gin.Context) {
	var (
		tr   request.TaskRequest
		resp response.TaskResponse
	)
	taskID, err := strconv.ParseUint(c.Param("taskID"), 10, 32)
	userID := c.GetUint("userID")

	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := c.ShouldBind(&tr); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}

	t, err := database.QueryTaskByTaskID(uint(taskID))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("找不到任务"))
		return
	}
	if t.OwnerID != userID && utils.CheckUserInGroup(userID, t.GroupID) <= model.RoleMember {
		c.JSON(http.StatusUnauthorized, response.MakeFailedResponse("没有修改这个任务的权限"))
		return
	}

	if err := copier.Copy(&t, &tr); err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("结构转换错误"))
		return
	}
	if err := database.TaskUpdate(&t); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("更新失败 "+err.Error()))
		return
	}
	if err := copier.Copy(&resp, &t); err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("结构转换错误"))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(gin.H{"task": resp}))
}

// GET: "/api/tasks"
// 查询任务
func GetTasksHandler(c *gin.Context) {
	userID := c.GetUint("userID")
	targetGroupsID := make([]uint, 0)
	joinedGroupsID := make([]uint, 0)
	joinedGroups, err := database.FindUserJoinedGroups(userID)
	// 获取已经加入的群组
	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	for _, g := range *joinedGroups {
		joinedGroupsID = append(joinedGroupsID, g.GroupID)
	}

	// 获取想要查询的群组
	targetGroups := c.QueryArray("groupID")
	for _, s := range targetGroups {
		if id, err := strconv.ParseUint(s, 10, 32); err != nil {
			c.JSON(http.StatusBadRequest, response.InvalidInfoError)
			return
		} else {
			targetGroupsID = append(targetGroupsID, uint(id))
		}
	}

	// 获取查询关键词
	word := c.Query("word")

	tasks, err := database.QueryTasks(joinedGroupsID, targetGroupsID, word)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse(err.Error()))
		return
	}
	// 转换为返回形式
	resp := make([]response.TaskResponse, 0)
	for _, t := range tasks {
		var r response.TaskResponse
		if err := copier.Copy(&r, &t); err != nil {
			c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("结构转换错误"))
			return
		}
		resp = append(resp, r)
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(gin.H{"tasks": resp}))
}
