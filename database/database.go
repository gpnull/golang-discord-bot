package database

import (
	"log"

	"github.com/gpnull/golang-github.com/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(connectionString string) {
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		// Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.TimekeepingStatus{},
		&models.TimekeepingExtraStatus{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Failed to close database connection: %v", err)
	}
	sqlDB.Close()
}
