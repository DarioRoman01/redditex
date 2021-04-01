package db

import (
	"fmt"
	"lireddit/env"

	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connect the server to postgres
func Connect() (*gorm.DB, error) {
	if err := cleanenv.ReadEnv(&env.Cfg); err != nil {
		return nil, fmt.Errorf("unable to connect to postgres: %s", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.Cfg.DBHost, env.Cfg.DBUser, env.Cfg.DBPassword, env.Cfg.DBName, env.Cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to postggres: %s", err)
	}

	return db, nil
}
