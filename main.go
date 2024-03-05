package main

import (
	"database/sql"
	"fiberapp/routers"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB)

	db, _ := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	routers.SetupRoutes(app)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Listen(":3000")
}
