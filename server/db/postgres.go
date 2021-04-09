package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// connect the server to postgres
func Connect() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to read env")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{Logger: newLogger})

	if err != nil {
		return nil, fmt.Errorf("unable to connect to postggres: %s", err)
	}

	return db, nil
}
