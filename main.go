package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=postgres password=marcel dbname=learn_fiber_gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	app := fiber.New()

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	users := app.Group("/users")

	users.Get("/", func(c *fiber.Ctx) error {
		users := new(User)
		db.Find(&users, "")

		return c.JSON(users)
	})

	users.Post("/", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			return err
		}

		log.Println(user.LastName)
		log.Println(user.FirstName)

		db.Create(&user)

		return c.SendString("Post: /users/")
	})

	users.Put("/", func(c *fiber.Ctx) error {
		return c.SendString("Put: /users/")
	})

	users.Delete("/", func(c *fiber.Ctx) error {
		return c.SendString("Delete: /users/")
	})

	app.Listen(":3000")
}
