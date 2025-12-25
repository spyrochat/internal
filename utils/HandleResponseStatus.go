package utils

import "github.com/gofiber/fiber/v2"

const (
   STATUS_SUCCESS string = "success"
   STATUS_ERROR   string = "error"
)

func SuccessMessage(c *fiber.Ctx, message string, code int, data any) error {
   if code == 0 {
      code = fiber.StatusOK
   }

   res := fiber.Map{
      "code": code,
      "status":  STATUS_SUCCESS,
      "message": message,
   }

   if data != nil {
      res["data"] = data
   }

   return c.JSON(res)
}

func ErrorMessage(c *fiber.Ctx, message string, code int) error {
   if code == 0 {
      code = fiber.StatusInternalServerError
   }

   return c.JSON(fiber.Map{
      "code": code,
      "status":  STATUS_ERROR,
      "message": message,
   })
}