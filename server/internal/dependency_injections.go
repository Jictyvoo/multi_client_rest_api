package internal

import (
	"errors"
	"fmt"
	"github.com/jictyvoo/multi_client_rest_api/modules"
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
		remy.Instance(func(retriever remy.DependencyRetriever) modules.DatabaseConfigABZ {
			tempConf := conf.Database.Abz1
			return modules.DatabaseConfigABZ{
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
		remy.Instance(func(retriever remy.DependencyRetriever) modules.DatabaseConfigXYC {
			tempConf := conf.Database.Xyc2
			return modules.DatabaseConfigXYC{
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
	modules.BindInjections(abzInjector)
	remy.Register(
		injector,
		remy.Factory(func(remy.DependencyRetriever) services.ContactsServiceFacade {
			return modules.ABZContactsServiceFactory.New(abzInjector)
		}),
		utils.ServiceABZ1,
	)

	// Bind the XYC_2 Service
	xycInjector := remy.NewInjector(remy.Config{GenerifyInterfaces: false, ParentInjector: injector})
	modules.BindInjections(xycInjector)
	remy.Register(
		injector,
		remy.Factory(func(retriever remy.DependencyRetriever) services.ContactsServiceFacade {
			return modules.XYCContactsServiceFactory.New(xycInjector)
		}),
		utils.ServiceXYC2,
	)
	return
}
