package utils

import "github.com/fasthttp/websocket"

type Client struct {
	Conn *websocket.Conn
}
