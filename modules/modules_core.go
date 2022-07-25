package modules

import (
	abzEntities "github.com/jictyvoo/multi_client_rest_api/modules/apps/abzcore/domain/entities"
	abzInfra "github.com/jictyvoo/multi_client_rest_api/modules/apps/abzcore/infra"
	xycEntities "github.com/jictyvoo/multi_client_rest_api/modules/apps/xyccore/domain/entities"
	xycInfra "github.com/jictyvoo/multi_client_rest_api/modules/apps/xyccore/infra"
	"github.com/jictyvoo/multi_client_rest_api/modules/internal/domain"
	"github.com/jictyvoo/multi_client_rest_api/modules/internal/domain/entities"
	"github.com/jictyvoo/multi_client_rest_api/modules/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/internal/infra"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/services"
	"github.com/wrapped-owls/goremy-di/remy"
)

type (
	DatabaseConfigXYC = xycInfra.DatabaseConfig
	DatabaseConfigABZ = abzInfra.DatabaseConfig
)

func BindInjections(injector remy.Injector) {
	infra.RegisterBinds(injector)
}

type ContactsServiceFactory[T entities.ContactEntity] struct {
}

func (serviceFactory ContactsServiceFactory[T]) New(injector remy.Injector) services.ContactsServiceFacade {
	return domain.NewContactsService[T](
		remy.Get[interfaces.ContactsRepository](injector),
	)
}

// Export service factories
var (
	ABZContactsServiceFactory = ContactsServiceFactory[*abzEntities.ContactEntity]{}
	XYCContactsServiceFactory = ContactsServiceFactory[*xycEntities.ContactEntity]{}
)
