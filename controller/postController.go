package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinheehanaaa/go-blog-backend/database"
	"github.com/jinheehanaaa/go-blog-backend/models"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse body")
	}
	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid paylod",
		})
	}
	return c.JSON(fiber.Map{
		"message": "You created a new post!",
	})
}

/*
func AllPost(c *fiber.Ctx) error {

}
*/
