package src

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"example.com/accounting/src/db"
	"example.com/accounting/src/routes"
	"example.com/accounting/src/services"
	"example.com/accounting/src/services/auth"
	"example.com/accounting/src/services/websocket"
)

func NewServer() *gin.Engine {


	mediaServerSockets := make(map[string]websocket.MediaServer) 
	
	router := gin.New()
	DB := db.MakeDB()
	service := services.MakeService(DB)
	controller := routes.Controller{Service: service, MediaServerSockets: mediaServerSockets}
	// var envVars = utils.GetEnvVars()

	// if envVars.DebugMode {
	// 	gin.SetMode(gin.DebugMode)
	// } else {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	// middlewares

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.Default())
	// router.Use(Auth())

	// static files serving
	router.Static("/images", "./images")

	auth := router.Group("auth")
	{
		v1 := auth.Group("v1")
		{
			controllers := controller
			v1.POST("/signup", controllers.SignUp)
			v1.POST("/login", controllers.Login)
			v1.GET("/getUser/:email", Auth(), controllers.GetUser)
			
			// news.GET("webRTCConnect")
		}
	}
	connect := router.Group("connect")
	{
		v1 := connect.Group("v1")
		{
			v1.GET("mediaServer/:mediaServerName", controller.ConnectWebSocket)
			v1.POST("client/:mediaServerName", controller.ClientConnect)
		}
	}
	// description string, mediaServerName

	return router
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Set("Email", claims.Email)
		context.Set("User", claims.Username)
		context.Next()
	}
}
