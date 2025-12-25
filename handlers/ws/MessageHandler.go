package ws

import (
   "fmt"
   "github.com/gofiber/contrib/websocket"
)

func MessageHandler(c *websocket.Conn) {
   userId := c.Params("id")
	fmt.Printf("User %s connected!\n", userId)

   for {
		// Read message from browser
		mt, msg, err := c.ReadMessage()
		if err != nil {
			break // End connection if error (user closed tab)
		}
		
		fmt.Printf("Received: %s\n", msg)

		// Send message back to browser (Echo)
		if err := c.WriteMessage(mt, msg); err != nil {
			break
		}
	}
}
