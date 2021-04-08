package db

import (
	"fmt"
	"lireddit/env"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// connect the server to postgres
func Connect() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	if err := cleanenv.ReadEnv(&env.Cfg); err != nil {
		return nil, fmt.Errorf("unable to connect to postgres: %s", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.Cfg.DBHost, env.Cfg.DBUser, env.Cfg.DBPassword, env.Cfg.DBName, env.Cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		return nil, fmt.Errorf("unable to connect to postggres: %s", err)
	}

	return db, nil
}
