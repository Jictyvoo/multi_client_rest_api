package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/dtos"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/corerrs"
	apiDtos "github.com/jictyvoo/multi_client_rest_api/services/apicontracts/dtos"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/services"
	"github.com/wrapped-owls/goremy-di/remy"
)

type ContactsController struct {
	injector remy.Injector
}

func NewContactsController(injector remy.Injector) *ContactsController {
	return &ContactsController{injector: injector}
}

func (ctrl *ContactsController) Bind(router fiber.Router) {
	router.Post("/", ctrl.Insert)
}

func (ctrl ContactsController) Insert(c *fiber.Ctx) error {
	var contactsList dtos.ContactsListDTO
	err := c.BodyParser(&contactsList)
	if err != nil {
		return err
	}

	// Check if the key in context
	serviceName := c.Locals("service-name", "").(string)
	service := remy.Get[services.ContactsServiceFacade](ctrl.injector, serviceName)
	if service == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"content": "Unable to access service for given client"},
		)
	}

	// Use the service to validate fields and execute the add function
	contactsDtoList := make([]apiDtos.ContactsDTO, 0, len(contactsList.Contacts))
	for _, contact := range contactsList.Contacts {
		contactsDtoList = append(contactsDtoList, contact)
	}
	err = service.Add(contactsDtoList)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if err == corerrs.ErrContactAlreadyExists {
			statusCode = fiber.StatusConflict
		} else if err == corerrs.ErrInvalidPhone {
			statusCode = fiber.StatusBadRequest
		}
		return c.Status(statusCode).JSON(
			fiber.Map{"content": "Unable to add contacts", "error": err.Error()},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{"content": "Contacts added successfully"},
	)
}
