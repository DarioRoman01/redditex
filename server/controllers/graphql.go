package controllers

import (
	"context"
	"lireddit/graph/generated"
	"lireddit/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Ctx context.Context
}

// insert echo request context in the context
func Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "EchoContextKey", c)
		c.SetRequest(c.Request().WithContext(ctx))
		cc := &CustomContext{c, ctx}
		return next(cc)
	}
}

// handler for graphql request
func GraphqlHandler(c echo.Context) error {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolvers.Resolver{},
			},
		),
	)

	cc := c.(*CustomContext)
	req := cc.Request()
	res := cc.Response()
	h.ServeHTTP(res, req)
	return nil
}

// handler for playground
func PlaygroundHandler(c echo.Context) error {
	ph := playground.Handler("GraphQL", "/graphql")
	cc := c.(*CustomContext)
	req := cc.Request()
	res := cc.Response()
	ph.ServeHTTP(res, req)
	return nil
}
