package authcore

import (
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/dtos"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/interfaces"
	"github.com/wrapped-owls/goremy-di/remy"
)

func CreateCustomer(injector remy.Injector, name string, namespace string, key string) error {
	service := domain.NewCustomerAuthService(
		remy.Get[interfaces.CustomersRepository](injector),
		remy.Get[string](injector, "security.secret_key"),
	)
	return service.Register(name, namespace, key)
}

func CustomerFindChecker(injector remy.Injector) func(string) (*dtos.CustomerDTO, error) {
	repo := remy.Get[interfaces.CustomersRepository](injector)
	return func(name string) (*dtos.CustomerDTO, error) {
		customer, err := repo.FindByName(name)
		if err != nil {
			return nil, err
		}
		return &customer, nil
	}
}
