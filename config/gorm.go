package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/okyanawang/money-transfer-go/httpserver/repository/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresGORM() (*gorm.DB, error) {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	pass := os.Getenv("PGPASSWORD")
	name := os.Getenv("PGDATABASE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Errorf("Failed to create uuid-ossp extension: %v", err)
		return nil, err
	}

	db.Debug().AutoMigrate(
		models.Account{},
		models.Transaction{},
		models.Callback{},
	)

	return db, nil
}
