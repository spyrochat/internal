package utils

import (
	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Conn *websocket.Conn
}

// --- Map of all active connections ---
var ActiveUsers = make(map[string]*Client)

// --- A channel to send messages into ---
var Broadcast = make(chan [2]string)
