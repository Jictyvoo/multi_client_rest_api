package internal

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recovery "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/config"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/controllers"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/utils"
	"github.com/jictyvoo/multi_client_rest_api/services/apicontracts/services"
	"github.com/wrapped-owls/goremy-di/remy"
	"log"
	"runtime"
)

const defaultStackTraceLength = 1024

func CatchStackTrace() []byte {
	buf := make([]byte, defaultStackTraceLength)
	bytesWritten := runtime.Stack(buf, false)
	return buf[:bytesWritten]
}

func bindInjections() (injector remy.Injector, err error) {
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

	// Bind the ABZ_1 Service
	remy.Register(
		injector,
		remy.Factory(func(retriever remy.DependencyRetriever) services.ContactsServiceFacade {
			// TODO: Implement the factory
			return nil
		}),
		utils.ServiceABZ1,
	)

	// Bind the XYC_2 Service
	remy.Register(
		injector,
		remy.Factory(func(retriever remy.DependencyRetriever) services.ContactsServiceFacade {
			// TODO: Implement the factory
			return nil
		}),
		utils.ServiceXYC2,
	)
	return
}

func SetupApp(data config.AppConfig, closeServerChan chan string) *fiber.App {
	// start bind the injections
	injector, err := bindInjections()
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

	controllers.NewContactsController(injector).
		Bind(app.Group("/contacts"))
	return app
}
