package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/accounting/src/routes/validators"
	"example.com/accounting/src/services"
	"github.com/gorilla/websocket"
)

// NewsController <controller>
// is used for describing controller actions for news.
type Controller struct {
	Service services.Service
	MediaServerSockets map[string]*websocket.Conn
}

// Get <function>
// is used to handle get action of news controller which will return <count> number of news.
// url: /v1/news?count=80 , by default <count> = 50
func (nc Controller) Login(c *gin.Context) {
	loginParam := validators.LoginParam{}
	err := c.BindJSON(&loginParam)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := nc.Service.Login(loginParam)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

// GetSources <function>
// is used to handle get action of news controller which will return all news sources.
// url: /v1/news/sources
func (nc Controller) SignUp(c *gin.Context) {
	createUser := validators.CreateUser{}
	err := c.BindJSON(&createUser)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := nc.Service.CreateUser(createUser)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(200, gin.H{
		"method":  user,
		"message": "Hello from GetSources function!",
	})
}

// GetTypes <function>
// is used to handle get action of news controller which will return all news types.
// url: /v1/news/types
func (nc Controller) GetUser(c *gin.Context) {
	email := c.Param("email")
	user, err := nc.Service.GetUser(email)

	if err != nil {
		c.JSON(404, "hi")
		return
	}

	c.JSON(200, gin.H{
		"method":  user,
		"message": "Hello from GetSources function!",
	})
}

func (nc Controller) ConnectWebSocket(c *gin.Context) {
	mediaServerName := c.Param("mediaServerName")
	services.WebSocketHandler(c, nc.MediaServerSockets, mediaServerName)
}

func (nc Controller) ClientConnect(c *gin.Context) {
	mediaServerName := c.Param("mediaServerName")
	clientConnectParams := validators.ClientConnectParam{}
	err := c.BindJSON(&clientConnectParams)

	if err != nil {
		c.JSON(404, "hi")
		return
	}

	answer, err := services.ClientConnect(c, nc.MediaServerSockets, mediaServerName, clientConnectParams.Description)
	println("done answer")
	println(answer)
	if err != nil {
		c.JSON(404, "Could Not find given Media Server")
		return
	}

	c.JSON(200, gin.H{
		"answer": answer,
	})
}
