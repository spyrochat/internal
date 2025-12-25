package main

import (
   "spyrochat/routes"
   "spyrochat/middlewares"

   "os"
   "log"
   "path/filepath"
   "github.com/gofiber/fiber/v2"
)

func main() {
   cwd, err := os.Getwd()
   if err != nil { log.Fatal(err.Error()) }
   var _ = filepath.Join(cwd, "..", "spyrochat.sock")

   //os.Remove(SOCK_FILE)

   //ln, err := net.Listen("unix", SOCK_FILE)
   if err != nil { panic(err.Error()) }

   //os.Chmod(SOCK_FILE, 0777)

   //defer ln.Close()
	//defer os.Remove(SOCK_FILE)
	
	app := fiber.New()
	
	wsRoute := app.Use("/ws", middlewares.WebsocketMiddleware)
   authenticatedRoute := app.Use(middlewares.AuthMiddleware)

   routes.WsRouter(wsRoute)
   routes.Router(authenticatedRoute)

	log.Fatal(app.Listen(":8000"))
}
