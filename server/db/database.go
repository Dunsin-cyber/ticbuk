package db

import (
	"fmt"

	"github.com/Dunsin-cyber/ticbuk/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigrator func(*gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		echo.New().Logger.Fatal("Failed to connect to database:", err)
	}

	log.Info("connected to the database")

	if err := DBMigrator(db); err != nil {
		echo.New().Logger.Fatal("Database migration failed:", err)
	}

	log.Info("DB migration successful")

	return db
}
