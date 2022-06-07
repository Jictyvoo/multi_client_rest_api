package authware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/memory"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/dtos"
	"time"
)

const (
	cacheDuration         = 30 * time.Minute
	DefaultContextKey     = "secret-user-data"
	DefaultAuthContextKey = "auth-token"
)

// Config defines the config for middleware.
type (
	TokenChecker func(authToken string) (*dtos.CustomerDTO, error)

	// Config defines all configs needed by the middleware
	Config struct {
		// Next defines a function to skip this middleware when returned true.
		//
		// Optional. Default: nil
		Next func(context *fiber.Ctx) bool

		// NamespaceChecker defines a function that checks user and returns its ID, in case user not found, return 0
		//
		// Required. Default: nil
		NamespaceChecker TokenChecker

		// CheckerError defines the response body for unauthorized responses.
		// By default, it will return with a 401 CheckerError
		//
		// Optional. Default: nil
		CheckerError fiber.Handler

		// Forbidden defines the response body for unauthorized responses.
		// By default it will return with a 403 Forbidden
		//
		// Optional. Default: nil
		Forbidden fiber.Handler

		// AuthContextKey is the key to look out the auth name in Locals
		//
		// Optional. Default: "auth-token"
		AuthContextKey string

		// NamespaceContextKey is the key to store the userID in Locals
		//
		// Optional. Default: "secret-user-id"
		NamespaceContextKey string

		// CacheStorage is the storage used to store users in cache, to preventing searching them multiple times
		//
		// Optional. Default: memory
		CacheStorage fiber.Storage
	}
)

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:                nil,
	Forbidden:           nil,
	AuthContextKey:      DefaultAuthContextKey,
	NamespaceContextKey: DefaultContextKey,
	CacheStorage:        memory.New(),
}

// Helper function to set default values
func configDefault(authConfigs ...Config) Config {
	// Return default authConfigs if nothing provided
	if len(authConfigs) < 1 {
		return ConfigDefault
	}

	// Override default authConfigs
	config := authConfigs[0]

	// Set default values
	if config.Next == nil {
		config.Next = ConfigDefault.Next
	}
	if config.NamespaceChecker == nil {
		panic("A NamespaceChecker function is required")
	}
	if config.CheckerError == nil {
		config.CheckerError = func(context *fiber.Ctx) error {
			return context.Status(fiber.StatusUnauthorized).
				SendString("User token not corresponds with any user founded.")
		}
	}
	if config.Forbidden == nil {
		config.Forbidden = func(context *fiber.Ctx) error {
			return context.Status(fiber.StatusForbidden).
				SendString("Forbidden app, your user has no permission on this route.")
		}
	}
	if config.AuthContextKey == "" {
		config.AuthContextKey = ConfigDefault.AuthContextKey
	}
	if config.NamespaceContextKey == "" {
		config.NamespaceContextKey = ConfigDefault.NamespaceContextKey
	}
	if config.CacheStorage == nil {
		config.CacheStorage = ConfigDefault.CacheStorage
	}
	return config
}
