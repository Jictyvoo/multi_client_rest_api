package interfaces

import "github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/dtos"

type CustomersRepository interface {
	FindByName(name string) (customer dtos.CustomerDTO, err error)
}
