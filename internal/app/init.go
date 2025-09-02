package app

import (
	"Educational-API-DBeaver-Sample-Database/internal/repository"
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func init_SQLiteDB() (*sql.DB, error) {
	logrus.Info("Database initialization started")

	dbName := viper.GetString("db.name")
	dbPath := os.Getenv("DBPATHE")

	logrus.Debug("dbPath: ", dbPath)
	if dbName == "" || dbPath == "" {
		return nil, fmt.Errorf("database path missing")
	}

	cfg := &repository.ConfigSQLite{
		DBPath:      dbPath + dbName,
		CacheSize:   viper.GetInt("db.cache_size"),
		JournalMode: viper.GetString("db.journal_mode"),
	}

	db, err := repository.NewSQLiteDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize db: %w", err)
	}

	logrus.Print("Database initialization completed")
	return db, nil
}
