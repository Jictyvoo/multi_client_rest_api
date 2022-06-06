package utils

import "github.com/gofiber/fiber/v2"

type (
	RouteBinder interface {
		Bind(fiber.Router)
	}

	// ContactsServiceFacade defines how the contacts service facade should be implemented and methods that it have to
	// manipulate contacts
	ContactsServiceFacade interface {
		Validate() error
	}
)
