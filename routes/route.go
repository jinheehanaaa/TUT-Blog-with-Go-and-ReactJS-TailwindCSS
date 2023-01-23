package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinheehanaaa/go-blog-backend/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
