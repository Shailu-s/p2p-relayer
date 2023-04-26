package db

import (
	"github.com/sirupsen/logrus"
	"github.com/volmexfinance/relayers/relayer-srv/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

// Config sqlite database configuration
type Config struct {
	FileName string
}

// DataBase is a struct for work with postgres DataBase via gorm
type DataBase struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

// NewDataBase creates a new DataBase instance
func NewDataBase(logger *logrus.Logger, config Config) (*DataBase, error) {
	// TODO: Hook gorm logger with logrus logger
	gormConfig := &gorm.Config{Logger: gorm_logger.Default.LogMode(gorm_logger.Silent)}

	db, err := gorm.Open(sqlite.Open(config.FileName), gormConfig)
	if err != nil {
		return nil, err
	}

	return &DataBase{DB: db, Logger: logger}, nil
}

// InitialMigration initialises the SQLite database schema
func (db *DataBase) InitialMigration() error {
	if err := db.DB.AutoMigrate(&models.Order{}, &models.Assets{}, &models.TransactionSent{}, &models.TransactionLog{}); err != nil {
		return err
	}

	return nil
}
