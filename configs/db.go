package configs

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dsn string) (*gorm.DB,error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("error when connecting db : ",err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)                  // koneksi idle
	sqlDB.SetMaxOpenConns(100)                // max koneksi aktif
	sqlDB.SetConnMaxLifetime(time.Hour)       // umur koneksi
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	log.Println("DB connected successfully")
	return db, nil

}
