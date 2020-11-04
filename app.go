package main

import (
	"backend/config"
	"github.com/gofiber/fiber"
	"log"
)

func main() {
	app := fiber.New()
	err := config.Configure(app)
	if err != nil {
		panic(err)
		return
	}

	log.Fatal(app.Listen(":3000"))
}
