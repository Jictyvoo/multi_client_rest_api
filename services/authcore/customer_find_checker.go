package authcore

import (
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/dtos"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/interfaces"
	"github.com/wrapped-owls/goremy-di/remy"
)

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
