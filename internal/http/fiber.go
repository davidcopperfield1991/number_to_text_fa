package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	ntt "main.go/internal/NTT"
)

type Number struct {
	Num int
}

func Init() {
	app := fiber.New()

	app.Post("/", numtotext)

	log.Fatal(app.Listen(":5533"))
}

// Handler
func numtotext(c *fiber.Ctx) error {
	var num Number
	c.BodyParser(&num)
	l := ntt.Call(num.Num)
	fmt.Println(ntt.Call(num.Num))
	return c.JSON(l)
}
