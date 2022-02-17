package app

import (
	"barlau/app/templates"
	"github.com/gofiber/fiber/v2"
)

type Manifest struct {
	Url  string `json:"url"`
	Type bool   `json:"native"`
}

func createPlayer(ctx *fiber.Ctx) error {
	m := new(Manifest)

	if err := ctx.BodyParser(m); err != nil {
		return err
	}

	var path = templates.CreateSamplePlayer(m.Url, m.Type)

	err := ctx.JSON(&fiber.Map{
		"url":    m.Url,
		"native": m.Type,
		"path":   path,
	})
	if err != nil {
		return err
	}
	return nil
}
