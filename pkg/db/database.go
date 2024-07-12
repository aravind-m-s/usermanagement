package db

import (
	"fmt"
	"usermanagement/pkg/config"
	"usermanagement/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb(cnf config.Config) (*gorm.DB, error) {
	dbInfo := fmt.Sprintf("host=%s users=%s dbname=%s port=%s password=%s", cnf.DBHost, cnf.DBUser, cnf.DBName, cnf.DBPort, cnf.DBPassword)
	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{SkipDefaultTransaction: true})

	db.AutoMigrate(&domain.Users{})

	return db, err
}
