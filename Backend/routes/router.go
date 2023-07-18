package routes

import "github.com/gin-gonic/gin"

func SetupRoutes() {

	e := gin.Default()

	registerUserRoutes(e)
	registerGroupsRoutes(e)
	registerPassWordRoutes(e)
	registerTasksRoutes(e)

}
