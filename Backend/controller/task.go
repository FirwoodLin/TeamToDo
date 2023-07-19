package controller

import (
	"TeamToDo/database"
	"TeamToDo/model"
	"TeamToDo/model/request"
	"TeamToDo/model/response"
	"TeamToDo/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// POST: "/api/tasks"
// 新建任务
func CreateTaskHandler(c *gin.Context) {
	var (
		tr request.TaskRequest
		t  model.Task
	)
	userID := c.GetUint("userID")
	groupID, err := strconv.ParseUint(c.PostForm("groupID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, response.InvalidInfoError)
		return
	}
	t.GroupID = uint(groupID)
	t.OwnerID = userID
	if utils.CheckUserInGroup(userID, t.GroupID) == model.RoleVisitor {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("你并不在该群组中"))
		return
	}

	if err := copier.Copy(&t, &tr); err != nil {
		c.JSON(http.StatusInternalServerError, response.MakeFailedResponse("结构转换错误"))
		return
	}

	// 在这里看看copy的效果
	log.Println(t)

	if err := database.TaskCreate(&t); err != nil {
		c.JSON(http.StatusBadRequest, response.MakeFailedResponse("创建任务失败 "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.MakeSucceedResponse(""))
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
		c.Abort()
		return
	}
	// if err := database.
	// 删除任务，未完待续
}

// PUT: "/api/tasks/:taskID"
// 修改任务
func UpdateTaskHandler(c *gin.Context) {
	var tr request.TaskRequest
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
		c.Abort()
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
	c.JSON(http.StatusOK, response.MakeSucceedResponse(map[string]interface{}{"task": t}))
}

// GET: "/api/tasks"
// 查询任务
// func GetTasksHandler(c *gin.Context) {
// 	userID := c.GetUint("userID")
// 	database.QueryTaskByTaskID()
// }
