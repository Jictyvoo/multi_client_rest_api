package abz_1_core

import (
	"database/sql"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/domain"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/infra/repositories"
	"github.com/wrapped-owls/goremy-di/remy"
)

func NewContactsService() *domain.ContactsService {
	return domain.NewContactsService(nil)
}

func NewContactsRepository(injector remy.Injector) interfaces.ContactsRepository {
	return repositories.NewContactsDbRepository(
		remy.Get[*sql.DB](injector),
	)
}
