package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recovery "github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/config"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/controllers"
	"github.com/jictyvoo/multi_client_rest_api/server/internal/utils"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore"
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
	jwtMiddleware := jwtware.New(jwtware.Config{
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
		Bind(app.Group("/contacts", jwtMiddleware))

	authcore.NewAuthController(injector).
		Bind(app.Group("/auth"))

	return app
}
