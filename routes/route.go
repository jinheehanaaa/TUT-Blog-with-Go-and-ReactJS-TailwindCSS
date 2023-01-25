package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinheehanaaa/go-blog-backend/controller"
	"github.com/jinheehanaaa/go-blog-backend/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticate) // Any API below this line can't access routes
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/allpost/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
}
