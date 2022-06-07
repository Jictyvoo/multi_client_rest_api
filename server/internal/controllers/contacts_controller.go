package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/dtos"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/utils"
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
	router.Get("/", ctrl.GetAll)
}

func (ctrl ContactsController) getService(c *fiber.Ctx) (services.ContactsServiceFacade, error) {
	// Check if the key in context
	var serviceName string
	undefinedTypeName := c.Locals("service-name")
	if undefinedTypeName != nil {
		serviceName = undefinedTypeName.(string)
	}
	if len(serviceName) == 0 {
		return nil, utils.ErrServiceNotFound
	}
	service := remy.Get[services.ContactsServiceFacade](ctrl.injector, serviceName)
	return service, nil
}

func (ctrl ContactsController) sendServiceError(c *fiber.Ctx, err error, service services.ContactsServiceFacade) (error, bool) {
	const message = "Unable to access service for given client"
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": err.Error(), "content": message},
		), true
	}
	if service == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"content": message},
		), true
	}
	return nil, false
}

func (ctrl ContactsController) Insert(c *fiber.Ctx) error {
	var contactsList dtos.ContactsListDTO
	if err := c.BodyParser(&contactsList); err != nil {
		return err
	}

	service, err := ctrl.getService(c)
	if sendErr, done := ctrl.sendServiceError(c, err, service); done || sendErr != nil {
		return sendErr
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

func (ctrl *ContactsController) GetAll(c *fiber.Ctx) error {
	service, err := ctrl.getService(c)
	if sendErr, done := ctrl.sendServiceError(c, err, service); done || sendErr != nil {
		return sendErr
	}

	var contactsList []apiDtos.ContactsDTO
	contactsList, err = service.ListAll()

	// TODO: Convert to presenter dto
	dtoList := dtos.ContactsListDTO{
		Contacts: make([]dtos.ContactsDTO, 0, len(contactsList)),
	}
	for _, contact := range contactsList {
		dtoList.Contacts = append(dtoList.Contacts, dtos.ContactsDTO{
			FullName:  contact.Name(),
			Cellphone: contact.Phone(),
		})
	}

	return c.JSON(dtoList)
}
