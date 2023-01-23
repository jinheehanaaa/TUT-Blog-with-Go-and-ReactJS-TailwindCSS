package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinheehanaaa/go-blog-backend/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	//app.Use(middleware.IsAuthenticate) // Any API below this line can't access routes
}
