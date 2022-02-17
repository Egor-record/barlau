package app

import (
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	serve(Routes(app))
}

func serve(router *fiber.App) {
	//addr := os.Getenv("SERVER_ADDRESS")
	//port := ":5000" // os.Getenv("SERVER_PORT")
	// router.Listen(fmt.Sprintf("%s:%s", addr, port))
	router.Listen("127.0.0.1:5000")
}
