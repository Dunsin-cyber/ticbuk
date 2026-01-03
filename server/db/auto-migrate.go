package db

import (
	"fmt"

	"github.com/Dunsin-cyber/ticbuk/models"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	// Auto migrate the Event model
	err := db.AutoMigrate(&models.Event{}, &models.Ticket{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate database: %v", err)
	}

	log.Info("Database auto-migration completed")
	return nil

}
