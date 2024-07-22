package models

import (
	"time"

	"gorm.io/gorm"
)

type TimekeepingExtraStatus struct {
	gorm.Model
	Day   string
	Start time.Time
	End   time.Time
}

func (TimekeepingExtraStatus) TableName() string {
	return "timekeeping_extra_status"
}
