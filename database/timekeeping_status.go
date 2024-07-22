package database

import (
	"fmt"

	"github.com/gpnull/golang-github.com/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Define the Database struct
type Database struct {
	DB *gorm.DB
}

// SaveTimeKeepingStatusButton saves or updates a timekeeping status button in SQL database
func (db *Database) SaveTimeKeepingStatusButton(timekeeping *models.TimekeepingStatus) error {
	result := db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(timekeeping)
	if result.Error != nil {
		return fmt.Errorf("failed to create/update button: %v", result.Error)
	}
	return nil
}

// GetTimeKeepingStatusButtons retrieves all timekeeping status buttons from SQL database
func (db *Database) GetTimeKeepingStatusButtons() ([]*models.TimekeepingStatus, error) {
	var buttons []*models.TimekeepingStatus
	result := db.DB.Find(&buttons)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve buttons: %v", result.Error)
	}
	return buttons, nil
}
