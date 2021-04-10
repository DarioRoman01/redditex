package main

import (
	"lireddit/cache"
	"lireddit/controllers"
	"lireddit/dataloaders"
	"lireddit/db"
	"lireddit/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	psql, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	psql.AutoMigrate(&models.User{}, &models.Post{}, &models.Updoot{})
	if err = godotenv.Load(); err != nil {
		log.Fatal("unable to read env")
	}
}

func main() {
	e := echo.New()
	store := cache.Client()

	// Middlewares
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("CORS_ORIGIN")},
		AllowCredentials: true,
	}))
	e.Use(controllers.Process)
	e.Use(cache.Sessions("qid", *store))
	e.Use(middleware.Recover())

	// Route => handler
	e.POST("/graphql", controllers.GraphqlHandler, dataloaders.DataLoaderMiddleware)
	e.GET("/graphql", controllers.PlaygroundHandler)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
