package internal

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recovery "github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/config"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/controllers"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/utils"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/services"
	"github.com/wrapped-owls/goremy-di/remy"
	"log"
	"runtime"
)

const (
	defaultStackTraceLength = 1024
	jwtContextKey           = "client-token"
)

func CatchStackTrace() []byte {
	buf := make([]byte, defaultStackTraceLength)
	bytesWritten := runtime.Stack(buf, false)
	return buf[:bytesWritten]
}

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

	injector = remy.NewInjector(remy.Config{GenerifyInterfaces: false})

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

	abzInjector := remy.NewInjector(remy.Config{GenerifyInterfaces: false, ParentInjector: injector})
	abz_1_core.BindInjections(abzInjector)

	// Bind the ABZ_1 Service
	remy.Register(
		injector,
		remy.Factory(func(remy.DependencyRetriever) services.ContactsServiceFacade {
			return abz_1_core.NewContactsService(abzInjector)
		}),
		utils.ServiceABZ1,
	)

	// Bind the XYC_2 Service
	remy.Register(
		injector,
		remy.Factory(func(retriever remy.DependencyRetriever) services.ContactsServiceFacade {
			return xyc_2_core.NewContactsService()
		}),
		utils.ServiceXYC2,
	)
	return
}

func SetupApp(data config.AppConfig, closeServerChan chan string) *fiber.App {
	// start bind the injections
	injector, err := bindInjections(data)
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New(fiber.Config{
		Prefork:      false,
		ErrorHandler: utils.DefaultErrorHandler,
		AppName:      "Multi-Client REST",
	})

	app.Use(
		recovery.New(recovery.Config{
			EnableStackTrace: true,
		}),
	)
	app.Use(logger.New())

	if closeServerChan != nil {
		go gracefulShutdown(app, closeServerChan)
	}

	// Create the JWT Middleware
	_ = jwtware.New(jwtware.Config{
		SigningKey:          []byte(data.Server.SymmetricKey),
		KeyRefreshInterval:  nil,
		KeyRefreshRateLimit: nil,
		SigningMethod:       jwtware.ES256,
		ContextKey:          jwtContextKey,
		TokenLookup:         "header:Authorization",
		AuthScheme:          "Bearer",
		KeyFunc:             nil,
	})

	controllers.NewContactsController(injector).
		Bind(app.Group("/contacts"))

	return app
}
