package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func checkPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	db, err := initDB()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer db.Close()

	api := app.Group("/tasks")

	api.Get("/", func(c *fiber.Ctx) error { return list(c, db) })
	api.Get("/:id", func(c *fiber.Ctx) error { return getTask(c, db) })
	api.Post("/", func(c *fiber.Ctx) error { return create(c, db) })
	api.Put("/:id", func(c *fiber.Ctx) error { return update(c, db) })
	api.Delete("/:id", func(c *fiber.Ctx) error { return remove(c, db) })

	log.Fatal(app.Listen(checkPort()))
}
