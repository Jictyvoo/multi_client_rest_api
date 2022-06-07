package infra

import (
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/infra/models"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/infra/repositories"
	"github.com/wrapped-owls/goremy-di/remy"
)

func RegisterRepositories(injector remy.Injector) {
	customersList := make([]models.CustomersModel, 0, 2)
	remy.Register(
		injector,
		remy.Factory(func(retriever remy.DependencyRetriever) interfaces.CustomersRepository {
			return repositories.NewCustomersRepositoryImpl(customersList)
		}),
	)
}
