package utils

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type errWrapper struct {
	Err     error           `json:"error"`
	Content json.RawMessage `json:"content"`
}

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	content := c.Response().Body()

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// check if the response is not a json, so wrap it as a json
	responseContentType := string(c.Response().Header.ContentType())
	if responseContentType != fiber.MIMEApplicationJSONCharsetUTF8 && responseContentType != fiber.MIMEApplicationJSON {
		var marshErr error
		content, marshErr = json.Marshal(string(content))
		if marshErr != nil {
			return marshErr
		}
	}

	responseBytes, marshErr := json.Marshal(&errWrapper{
		Err:     err,
		Content: content,
	})

	if marshErr != nil {
		return marshErr
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return c.Status(code).Send(responseBytes)
}
