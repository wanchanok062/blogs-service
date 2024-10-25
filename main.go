package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/wanchanok6698/web-blogs/api/v1/routes"
	"github.com/wanchanok6698/web-blogs/config"
)

func main() {
	config.ConnectDB()
	if config.DB == nil {
		log.Fatal("MongoDB connection failed")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load environment variables: ", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	app.Use(logger.New())

	app.Use(limiter.New(limiter.Config{
		Max:        15,
		Expiration: 30 * time.Second,
	}))

	prefix := os.Getenv("API_PREFIX")
	routes.BlogRoutes(prefix, app)

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
