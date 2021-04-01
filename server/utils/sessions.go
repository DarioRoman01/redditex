package utils

import (
	"context"
	"fmt"
	"lireddit/cache"
)

func GenerateSession(ctx context.Context, id int) error {
	ec, err := EchoContextFromContext(ctx)
	if err != nil {
		return fmt.Errorf("unable to get echo context")
	}

	session := cache.Default(ec)
	session.Set("userId", id)
	session.Save()
	return nil
}
