package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
)

var printBody = os.Getenv("PRINT_BODY") == "true"
var printHeaders = os.Getenv("PRINT_HEADERS") == "true"

func serverAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":5050"
	}

	return ":" + port
}

func handler(c *fiber.Ctx) error {
	rid := strings.Split(uuid.New().String(), "-")[0]

	fmt.Printf("[%s] %s %s\n", rid, c.Request().Header.Method(), c.Request().URI().String())

	if printBody {
		fmt.Printf("[%s] Start Body:\n", rid)
		fmt.Printf("%s", c.Body())
		fmt.Printf("\n[%s] End Body\n", rid)
	}

	if printHeaders {
		fmt.Printf("[%s] Start Headers:\n", rid)
		fmt.Printf("%s", c.Request().Header.Header())
		fmt.Printf("\n[%s] End Headers\n", rid)
	}

	return c.SendStatus(200)
}

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(cors.New())
	app.Use(handler)

	err := app.Listen(serverAddr())
	if err != nil {
		panic(err.Error())
	}
}
