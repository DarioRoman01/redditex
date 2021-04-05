package utils

import (
	"context"
	"fmt"
	"lireddit/cache"
)

// Generate session for users
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

// Get users session from the context
func GetUserSession(ctx context.Context) (int, error) {
	ec, err := EchoContextFromContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("cannot get echo context")
	}

	session := cache.Default(ec)
	val := session.Get("userId")
	if val == nil {
		return 0, fmt.Errorf("not authenticated")
	}
	userId := val.(int)
	return userId, nil
}
