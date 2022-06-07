package xyc_2_core

import (
	"database/sql"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/infra/repositories"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/services"
	"github.com/wrapped-owls/goremy-di/remy"
)

func NewContactsService() services.ContactsServiceFacade {
	return domain.NewContactsService(nil)
}

func NewContactsRepository(injector remy.Injector) interfaces.ContactsRepository {
	return repositories.NewContactsDbRepository(
		remy.Get[*sql.DB](injector),
	)
}
