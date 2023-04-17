package websocket

import (
	"github.com/gorilla/websocket"
)

type MediaServer struct {
	Connection *websocket.Conn
	Name       string
}
