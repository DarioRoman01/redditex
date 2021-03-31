package cache

import (
	"github.com/gorilla/sessions"
)

type CookieStore interface {
	Store
}

func NewCookieStore(keyPairs ...[]byte) CookieStore {
	return &cookieStore{sessions.NewCookieStore(keyPairs...)}
}

type cookieStore struct {
	*sessions.CookieStore
}

func (c *cookieStore) Options(options Options) {
	c.CookieStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}

func (c *cookieStore) MaxAge(age int) {
	c.CookieStore.MaxAge(age)
}
