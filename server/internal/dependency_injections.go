package internal

import (
	"errors"
	"fmt"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/config"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/utils"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/services"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore"
	"github.com/wrapped-owls/goremy-di/remy"
)

func bindInjections(conf config.AppConfig) (injector remy.Injector, err error) {
	defer func() {
		// if any panic is throw, then catches it and return as error
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
			buf := CatchStackTrace()

			// wrap the stacktrace in the error
			err = errors.New(fmt.Sprintf("%s\n%s", err, buf))
		}
	}()

	injector = remy.NewInjector(remy.Config{GenerifyInterfaces: false, UseReflectionType: true})

	remy.RegisterInstance(injector, conf.SymmetricKey, "security.secret_key")

	// Bind authcore service
	authcore.BindInjections(injector)

	remy.Register(
		injector,
		remy.Instance(func(retriever remy.DependencyRetriever) abz_1_core.DatabaseConfig {
			tempConf := conf.Database.Abz1
			return abz_1_core.DatabaseConfig{
				Host:     tempConf.Host,
				Port:     int(tempConf.Port),
				User:     tempConf.User,
				Password: tempConf.Password,
				Database: tempConf.Name,
			}
		}),
	)
	remy.Register(
		injector,
		remy.Instance(func(retriever remy.DependencyRetriever) xyc_2_core.DatabaseConfig {
			tempConf := conf.Database.Xyc2
			return xyc_2_core.DatabaseConfig{
				Host:     tempConf.Host,
				Port:     int(tempConf.Port),
				User:     tempConf.User,
				Password: tempConf.Password,
				Database: tempConf.Name,
			}
		}),
	)

	// Bind the ABZ_1 Service
	abzInjector := remy.NewInjector(remy.Config{GenerifyInterfaces: false, ParentInjector: injector})
	abz_1_core.BindInjections(abzInjector)
	remy.Register(
		injector,
		remy.Factory(func(remy.DependencyRetriever) services.ContactsServiceFacade {
			return abz_1_core.NewContactsService(abzInjector)
		}),
		utils.ServiceABZ1,
	)

	// Bind the XYC_2 Service
	xycInjector := remy.NewInjector(remy.Config{GenerifyInterfaces: false, ParentInjector: injector})
	xyc_2_core.BindInjections(xycInjector)
	remy.Register(
		injector,
		remy.Factory(func(retriever remy.DependencyRetriever) services.ContactsServiceFacade {
			return xyc_2_core.NewContactsService(xycInjector)
		}),
		utils.ServiceXYC2,
	)
	return
}
