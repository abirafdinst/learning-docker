package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		fmt.Println("Ouch ")
		return c.SendString("test")
	})

	log.Fatal(app.Listen(":77"))

}
