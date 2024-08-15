package purchasing

import "github.com/gofiber/fiber/v2"

func Purchase(c *fiber.Ctx) error {
	return c.SendString("Purchase")
}
