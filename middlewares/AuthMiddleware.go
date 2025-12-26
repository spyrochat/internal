package middlewares

import (
	"spyrochat/utils"

	"os"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
   token := c.Get("Authorization")

   if len(token) < 7 || token[:7] != "Bearer " {
      return utils.ErrorMessage(c, "Secret key not provided", fiber.StatusUnauthorized)
   }

   if token[7:] != os.Getenv("SPYROCHAT_UNIX_SECRET") {
      return utils.ErrorMessage(c, "Invalid or secret key", fiber.StatusUnauthorized)
   }

   return c.Next()
}