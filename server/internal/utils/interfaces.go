package utils

import "github.com/gofiber/fiber/v2"

type (
	RouteBinder interface {
		Bind(fiber.Router)
	}
)
