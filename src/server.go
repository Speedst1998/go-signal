package src

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"example.com/accounting/src/db"
	"example.com/accounting/src/routes"
	"example.com/accounting/src/services"
	"example.com/accounting/src/services/auth"
)

func NewServer() *gin.Engine {

	router := gin.New()
	DB := db.MakeDB()
	service := services.MakeService(DB)
	controller := routes.Controller{Service: service}
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

	v1 := router.Group("v1")
	{
		news := v1.Group("auth")
		{
			controllers := controller
			news.POST("/signup", controllers.SignUp)
			news.POST("/login", controllers.Login)
			news.GET("/getUser/:email", Auth(), controllers.GetUser)
		}
	}

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
