package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	e := gin.Default()

	e.Use(cors.Default())

	registerUserRoutes(e)
	registerGroupsRoutes(e)
	registerPassWordRoutes(e)
	registerTasksRoutes(e)
	return e
}
