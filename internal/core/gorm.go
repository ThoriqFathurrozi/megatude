package core

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ThoriqFathurrozi/megatude/configs"
	earthquake "github.com/ThoriqFathurrozi/megatude/internal/domains/earthquake/entity"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() (*gorm.DB, error) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	cfg := configs.GetConfig()

	user := cfg.Database.User
	password := cfg.Database.Password
	host := cfg.Database.Hostname
	port := cfg.Database.Port
	dbname := cfg.Database.Name

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Jakarta", host, user, password, dbname, port)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			Colorful:                  false,
			ParameterizedQueries:      true,
			IgnoreRecordNotFoundError: true,
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}),
		&gorm.Config{Logger: dbLogger})

	if err != nil {
		sugar.Error("Failed to connect to database", zap.Error(err))
	}

	if err := db.AutoMigrate(

		// Migrate Model
		earthquake.Earthquake{},
	); err != nil {
		sugar.Error("Failed to migrate database", zap.Error(err))
	}

	return db, nil
}
