package routes

import (
   "spyrochat/handlers"
   "github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
   router.Post("/notify", handlers.NotifyHandler)
}
