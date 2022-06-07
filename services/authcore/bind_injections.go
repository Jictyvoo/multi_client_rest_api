package authcore

import (
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/infra"
	"github.com/wrapped-owls/goremy-di/remy"
)

func BindInjections(injector remy.Injector) {
	infra.RegisterRepositories(injector)
}
