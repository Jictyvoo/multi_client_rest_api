package authcore

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/dtos"
	"github.com/wrapped-owls/goremy-di/remy"
)

type AuthController struct {
	service *domain.CustomerAuthService
}

func NewAuthController(injector remy.Injector) *AuthController {
	return &AuthController{service: domain.NewCustomerAuthService(nil, "")}
}

func (ctrl *AuthController) Bind(router fiber.Router) {
	router.Post("/access", ctrl.Access)
}

func (ctrl *AuthController) Access(c *fiber.Ctx) error {
	var customerDTO dtos.CustomerDTO
	if err := c.BodyParser(&customerDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	claims, err := ctrl.service.DoLogin(customerDTO.Name, customerDTO.Key)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"error":   err.Error(),
				"content": "Unable to login with given credentials",
			},
		)
	}

	var token string
	token, err = ctrl.service.CreateAccessToken(claims)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": token})
}
