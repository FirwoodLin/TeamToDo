package routes

import (
	"TeamToDo/controller"
	"TeamToDo/utils"

	"github.com/gin-gonic/gin"
)

func registerTasksRoutes(e *gin.Engine) {
	task := e.Group("/api/tasks", utils.MiddlewareJWTAuthorize())

	// 新建任务
	task.POST("/", controller.CreateTaskHandler)

	// 删除任务
	task.DELETE("/:taskID", controller.RemoveTaskHandler)

	// 修改任务
	task.PUT("/:taskID", controller.UpdateTaskHandler)

	// 查询任务
	// task.GET("/", controller.GetTasksHandler)
}
