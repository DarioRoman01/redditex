package utils

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

// Get echo from the graphql request context
func EchoContextFromContext(ctx context.Context) (echo.Context, error) {
	echoContext := ctx.Value("EchoContextKey")
	if echoContext == nil {
		err := fmt.Errorf("could not retrieve echo.Context")
		return nil, err
	}

	ec, ok := echoContext.(echo.Context)
	if !ok {
		err := fmt.Errorf("echo.Context has wrong type")
		return nil, err
	}
	return ec, nil
}
