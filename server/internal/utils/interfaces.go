package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/dtos"
)

type (
	RouteBinder interface {
		Bind(fiber.Router)
	}

	// ContactsServiceFacade defines how the contacts service facade should be implemented and methods that it have to
	// manipulate contacts
	ContactsServiceFacade interface {
		Add(contacts []dtos.ContactsDTO) error
	}
)
