package routes

import (
   "spyrochat/handlers/ws"

   "github.com/gofiber/fiber/v2"
   "github.com/gofiber/contrib/websocket"
)

func WsRouter(router fiber.Router) {
   router.Get("/:id", websocket.New(ws.MessageHandler))
}
