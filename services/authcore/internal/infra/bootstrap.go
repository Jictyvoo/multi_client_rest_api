package infra

import (
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/infra/models"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/infra/repositories"
	"github.com/wrapped-owls/goremy-di/remy"
)

func RegisterRepositories(injector remy.Injector) {
	remy.Register(
		injector,
		remy.LazySingleton(func(retriever remy.DependencyRetriever) *models.CustomersList {
			return &models.CustomersList{}
		}),
	)
	remy.Register(
		injector,
		remy.Factory(func(retriever remy.DependencyRetriever) interfaces.CustomersRepository {
			return repositories.NewCustomersRepositoryImpl(remy.Get[*models.CustomersList](retriever))
		}),
	)
}
