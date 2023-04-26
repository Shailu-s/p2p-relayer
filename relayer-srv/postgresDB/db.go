package postgresdb

import (
	"github.com/volmexfinance/relayers/relayer-srv/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// StorageConfig contains configurations for storage, postgreSQL
type PostgresStorageConfig struct {
	URL        string // DataBase URL for connection
	DBDriver   string // DataBase driver
	DBHOST     string // DataBase host
	DBPORT     int64
	DBSSL      string // DataBase sslmode
	DBName     string // DataBase name
	DBUser     string // DataBase's user
	DBPassword string // User's password
}

type PostgresDataBase struct {
	DB *gorm.DB
}

// TODO: Tx sent confirmed to be moved to postgres
func InitialMigration(dbURL string) (*PostgresDataBase, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Order{}, &models.Assets{})

	// sellOrder := storage.Order{
	// 	OrderID:      "ggggggggggggggggggg",
	// 	OrderType:    "0xf555eb98",
	// 	Trader:       "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
	// 	Deadline:     87654321987654,
	// 	Assets:       []storage.Assets{{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}, {VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}},
	// 	Price:        1,
	// 	Salt:         "44",
	// 	TriggerPrice: "0",
	// 	Sign:         "",
	// 	IsShort:      false,
	// 	Status:       storage.MatchedStatusFailedConfirmed,
	// }
	// sellOrder1 := storage.Order{
	// 	OrderID:      "KKKKKKKKK",
	// 	OrderType:    "0xf555eb98",
	// 	Trader:       "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
	// 	Deadline:     87654321987654,
	// 	Assets:       []storage.Assets{{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}, {VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}},
	// 	Price:        1,
	// 	Salt:         "44",
	// 	TriggerPrice: "0",
	// 	Sign:         "",
	// 	IsShort:      false,
	// 	Status:       storage.MatchedStatusFailedConfirmed,
	// }

	// db.Create(&sellOrder)
	// db.Create(&sellOrder1)
	return &PostgresDataBase{DB: db}, err
}
