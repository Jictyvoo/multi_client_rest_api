package xyc_2_core

import (
	"database/sql"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/infra"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/infra/repositories"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/services"
	"github.com/wrapped-owls/goremy-di/remy"
)

type (
	DatabaseConfig = infra.DatabaseConfig
)

func BindInjections(injector remy.Injector) {
	infra.RegisterDbConn(injector)
	infra.RegisterRepositories(injector)
}

func NewContactsService(injector remy.Injector) services.ContactsServiceFacade {
	repository := remy.Get[interfaces.ContactsRepository](injector)
	return domain.NewContactsService(repository)
}

func NewContactsRepository(injector remy.Injector) interfaces.ContactsRepository {
	return repositories.NewContactsDbRepository(
		remy.Get[*sql.DB](injector),
	)
}
