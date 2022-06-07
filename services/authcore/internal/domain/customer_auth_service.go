package domain

import (
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/dtos"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type CustomerAuthService struct {
	secretKey     string
	customersRepo interfaces.CustomersRepository
}

func NewCustomerAuthService(customersRepo interfaces.CustomersRepository, secretKey string) *CustomerAuthService {
	return &CustomerAuthService{customersRepo: customersRepo, secretKey: secretKey}
}

func (serv CustomerAuthService) DoLogin(name, key string) (utils.Claims, error) {
	customer, err := serv.customersRepo.FindByName(name)
	if err != nil {
		return utils.Claims{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(customer.Key), []byte(key)); err != nil {
		return utils.Claims{}, dtos.ErrInvalidCustomerKey
	}

	return utils.GenerateClaimToken(customer.Name), nil
}

func (serv CustomerAuthService) CreateAccessToken(claims utils.Claims) (string, error) {
	return utils.CreateJWT(claims, serv.secretKey)
}
