package main

import (
	"spyrochat/handlers/ws/utils"
	"spyrochat/middlewares"
	"spyrochat/routes"

	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func main() {
   cwd, err := os.Getwd()
   if err != nil { log.Fatal(err.Error()) }
   var _ = filepath.Join(cwd, "..", "spyrochat.sock")

   //os.Remove(SOCK_FILE)

   //ln, err := net.Listen("unix", SOCK_FILE)
   // if err != nil { panic(err.Error()) }

   //os.Chmod(SOCK_FILE, 0777)

   //defer ln.Close()
	//defer os.Remove(SOCK_FILE)

   // --- Create the worker pool ---
   const workerCount = 10

   for range workerCount {
      go utils.ReplyMessages()
   }

	app := fiber.New()

   authenticatedRoute := app.Group("/api", middlewares.AuthMiddleware)
	wsRoute := app.Group("/ws", middlewares.WebsocketMiddleware)

   routes.Router(authenticatedRoute)
   routes.WsRouter(wsRoute)

	log.Fatal(app.Listen(":8000"))
}
