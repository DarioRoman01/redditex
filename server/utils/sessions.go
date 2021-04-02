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

func GetUserSession(ctx context.Context) (interface{}, error) {
	ec, err := EchoContextFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot get echo context")
	}

	session := cache.Default(ec)
	val := session.Get("userId")
	if val == nil {
		return nil, fmt.Errorf("you are not logged in")
	}
	return val, nil
}
