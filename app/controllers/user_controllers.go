package controllers

import (
	"MarcelArt/learn_fiber_gorm/app/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitUserController(db *gorm.DB, app *fiber.App) {
	users := app.Group("/users")

	users.Get("/", func(c *fiber.Ctx) error {
		users := new(models.User)
		db.Find(&users, "")

		return c.JSON(users)
	})
}
