package database

import (
	"database/sql"
	"fmt"

	"github.com/baleegh-ud-din/hive/config"
	"github.com/baleegh-ud-din/hive/utils"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var logger = utils.NewLogger()

func Connect(cfg *config.Config) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		failmsg := fmt.Sprintf("❌ Failed to connect to database: %s", err)
		logger.Error(failmsg)
		return
	}

	err = DB.Ping()
	if err != nil {
		failmsg := fmt.Sprintf("❌ Failed to ping database: %s", err)
		logger.Error(failmsg)
		return
	}

	logger.Info("✅ Connected to Vendex Database")
	// CreateSchemas()
	// CreateMigrations()
}

func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			failmsg := fmt.Sprintf("❌ Failed to close database connection: %s", err)
			logger.Error(failmsg)
		} else {
			logger.Info("🔌 Database connection closed successfully")
		}
	}
}

func ValidateDB() {
	if DB == nil {
		logger.Error("🔌 Database connection is not initialized. Please check your Connect")
		return
	}
}
