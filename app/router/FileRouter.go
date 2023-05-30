package router

import (
	handlers "Gonverter/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterEndpoints(router fiber.Router) {
	router.Get("/status", func(context *fiber.Ctx) error {
		return context.JSON(fiber.Map{"response": "Ready"})
	})
	router.Post("/file/FromFile", handlers.SaveFile)
	router.Post("/file/FromBytes", handlers.SaveFileFromBytes)
	router.Post("/file/FromString", handlers.SaveFileFromString)
}
