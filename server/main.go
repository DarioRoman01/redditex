package main

import (
	"lireddit/cache"
	"lireddit/controllers"
	"lireddit/db"
	"lireddit/models"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	psql, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	psql.AutoMigrate(&models.User{}, &models.Post{}, &models.Updoot{})
}

func main() {
	e := echo.New()
	store, err := cache.NewRedisStore(32, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		log.Fatal("unable to connect to redis")
	}

	// Middlewares
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))
	e.Use(controllers.Process)
	e.Use(cache.Sessions("qid", store))
	e.Use(middleware.Recover())

	// Route => handler
	e.POST("/graphql", controllers.GraphqlHandler)
	e.GET("/graphql", controllers.PlaygroundHandler)

	// Start server
	e.Logger.Fatal(e.Start(":4000"))
}
