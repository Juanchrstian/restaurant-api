package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/juanchrstian/restaurant-api/internal/shared/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg config.DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	configurePool(sqlDB)

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func configurePool(db *sql.DB) {

	db.SetMaxOpenConns(25)

	db.SetMaxIdleConns(10)

	db.SetConnMaxLifetime(30 * time.Minute)

	db.SetConnMaxIdleTime(10 * time.Minute)

}