package repositories

import (
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/dtos"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/infra/models"
)

type CustomersRepositoryImpl struct {
	customersList []models.CustomersModel
}

func NewCustomersRepositoryImpl(customersList []models.CustomersModel) *CustomersRepositoryImpl {
	return &CustomersRepositoryImpl{customersList: customersList}
}

func (repo CustomersRepositoryImpl) FindByName(name string) (customer dtos.CustomerDTO, err error) {
	for _, c := range repo.customersList {
		if c.Name == name {
			return dtos.CustomerDTO{
				Name: c.Name,
				Key:  string(c.Key),
			}, nil
		}
	}

	err = dtos.ErrCustomerNotFound
	return
}
