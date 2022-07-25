package infra

import (
	abzInfra "github.com/jictyvoo/multi_client_rest_api/modules/apps/abzcore/infra"
	xycInfra "github.com/jictyvoo/multi_client_rest_api/modules/apps/xyccore/infra"
	"github.com/wrapped-owls/goremy-di/remy"
)

func RegisterBinds(injector remy.Injector) {
	// Register all injections from ABZ_1_CORE
	abzInfra.RegisterDbConn(injector)
	abzInfra.RegisterRepositories(injector)

	// Register all injections from XYC_2_CORE
	xycInfra.RegisterDbConn(injector)
	xycInfra.RegisterRepositories(injector)
}
