package ws

import (
	"spyrochat/handlers/ws/utils"
	"strings"

	"github.com/gofiber/contrib/websocket"
)

func MessageHandler(c *websocket.Conn) {
	client := &utils.Client{Conn: c}
	clientID := c.Params("id")

	utils.ActiveUsers[clientID] = client

	c.WriteMessage(1, []byte("Hello User " + clientID + "! You are now connected."))

	// --- Clean up when they disconnect ---
	defer func() {
		delete(utils.ActiveUsers, clientID)
		c.Close()
	}()

	// --- Listen for messages from this specific user ---
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			break // End connection if error (user closed tab)
		}

		parts := strings.SplitN(string(msg), ":", 2)
		var receiverID, senderMsg string

		if len(parts) == 2 {
			receiverID = parts[0]
			senderMsg = parts[1]
		} else {
			continue // Skip message if the ID is not provided
		}

		// --- Send the message to the "Broadcast" channel ---
		utils.Broadcast <- [2]string{receiverID, senderMsg}
	}
}
	