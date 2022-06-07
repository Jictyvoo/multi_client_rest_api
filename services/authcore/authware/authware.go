package authware

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jictyvoo/multi_client_rest_api/services/authcore/internal/domain/dtos"
)

// New creates a new middleware handler
func New(authConfig Config) fiber.Handler {
	// Set default authConfig
	config := configDefault(authConfig)

	// Return new handler
	return func(context *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if config.Next != nil && config.Next(context) {
			return context.Next()
		}

		// Get authorization header
		authToken, ok := context.Locals(config.AuthContextKey).(*jwt.Token)
		if ok && authToken != nil {
			// Decode the header contents

			// Get the user data
			var (
				userData = &dtos.CustomerDTO{}
				err      error
				data     []byte
			)

			var mapClaims jwt.MapClaims
			if mapClaims, ok = authToken.Claims.(jwt.MapClaims); ok {
				customerClaims := [...]string{
					"",
					"",
				}
				if mapClaims["name"] != nil {
					customerClaims[0] = mapClaims["name"].(string)
				}
				if mapClaims["Uuid"] != nil {
					customerClaims[1] = mapClaims["Uuid"].(string)
				}

				if data, err = config.CacheStorage.Get(customerClaims[1]); err == nil {
					err = json.Unmarshal(data, userData)
				}

				if err != nil || len(userData.Namespace) < 1 {
					userData, _ = config.NamespaceChecker(customerClaims[0])
					if len(userData.Namespace) < 1 {
						return config.CheckerError(context)
					}
					marshalData, _ := json.Marshal(userData)
					_ = config.CacheStorage.Set(customerClaims[1], marshalData, cacheDuration)
				}
				context.Locals(config.NamespaceContextKey, userData.Namespace)
				return context.Next()
			}
		}
		return config.Forbidden(context)
	}
}
