package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func InitDB(config Config) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName, config.DbSslMode,
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return err
	}
	log.Error("connected to db: " + config.DbName)

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(24 * time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	return err
}
func GetDBInstance() *gorm.DB {
	return db
}
