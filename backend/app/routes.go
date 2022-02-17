package app

import "github.com/gofiber/fiber/v2"

func Routes(router *fiber.App) *fiber.App {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	v1 := router.Group("/api/v1")

	v1.Post("/createPlayer", createPlayer)

	router.Static("/", "../frontend/build")

	return router
}
