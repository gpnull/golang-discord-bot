package database

import (
	"fmt"

	"github.com/gpnull/golang-github.com/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) SaveTimeKeepingStatusButton(timekeeping *models.TimekeepingStatus) error {
	result := db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(timekeeping)
	if result.Error != nil {
		return fmt.Errorf("db failed SaveTimeKeepingStatusButton: %v", result.Error)
	}
	return nil
}

func (db *Database) GetTimeKeepingStatusButtons() ([]*models.TimekeepingStatus, error) {
	var buttons []*models.TimekeepingStatus
	result := db.DB.Find(&buttons)
	if result.Error != nil {
		return nil, fmt.Errorf("db failed GetTimeKeepingStatusButtons: %v", result.Error)
	}
	return buttons, nil
}

func (db *Database) GetTimeKeepingStatusButtonByID(buttonID string) (*models.TimekeepingStatus, error) {
	var button *models.TimekeepingStatus
	result := db.DB.First(&button, "button_id = ?", buttonID)
	if result.Error != nil {
		return nil, fmt.Errorf("db failed GetTimeKeepingStatusButtonByID %s: %v", buttonID, result.Error)
	}
	return button, nil
}

// for OT

func (db *Database) SaveTimeKeepingOvertimeStatusButton(timekeepingOvertime *models.TimekeepingOvertimeStatus) error {
	result := db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(timekeepingOvertime)
	if result.Error != nil {
		return fmt.Errorf("db failed SaveTimeKeepingOvertimeStatusButton: %v", result.Error)
	}
	return nil
}

func (db *Database) GetTimeKeepingOvertimeStatusButtons() ([]*models.TimekeepingOvertimeStatus, error) {
	var buttons []*models.TimekeepingOvertimeStatus
	result := db.DB.Find(&buttons)
	if result.Error != nil {
		return nil, fmt.Errorf("db failed GetTimeKeepingOvertimeStatusButtons: %v", result.Error)
	}
	return buttons, nil
}

func (db *Database) GetTimeKeepingOvertimeStatusButtonByID(buttonID string) (*models.TimekeepingOvertimeStatus, error) {
	var button *models.TimekeepingOvertimeStatus
	result := db.DB.First(&button, "button_id = ?", buttonID)
	if result.Error != nil {
		return nil, fmt.Errorf("db failed GetTimeKeepingOvertimeStatusButtonByID %s: %v", buttonID, result.Error)
	}
	return button, nil
}
