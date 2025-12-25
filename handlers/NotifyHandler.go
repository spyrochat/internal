package handlers

import (
   "spyrochat/utils"

   "fmt"
   "github.com/gofiber/fiber/v2"
)

func NotifyHandler(c *fiber.Ctx) error {
   payload := make(map[string]interface{})

   if err := c.BodyParser(&payload); err != nil {
      return utils.ErrorMessage(c, "Invalid request payload", fiber.StatusUnprocessableEntity)
   }

   fmt.Println("Parsed User:", payload)

   return utils.SuccessMessage(c, "Go fiber received the data", 0, payload)
}