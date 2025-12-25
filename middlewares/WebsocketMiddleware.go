package middlewares

import (
   "github.com/gofiber/fiber/v2"
   "github.com/gofiber/contrib/websocket"
)

func WebsocketMiddleware(c *fiber.Ctx) error {
   if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}
	//return utils.ErrorMessage(c, "Invalid Request", fiber.ErrUpgradeRequired)
	return fiber.ErrUpgradeRequired
}
