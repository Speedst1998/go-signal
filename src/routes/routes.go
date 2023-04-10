package routes

import (
	"github.com/gin-gonic/gin"
)

// NewRouter <function>
// is used to create a GIN engine instance where all controller and routes will be placed
func MakeRouters(router *gin.Engine, controller Controller) *gin.Engine {

	// endpoints
	v1 := router.Group("v1")
	{
		news := v1.Group("auth")
		{
			controllers := controller

			news.GET("/getUser/:email", controllers.GetUser)
			news.POST("/signup", controllers.SignUp)
			news.GET("/login", controllers.Login)
		}
	}

	return router
}
