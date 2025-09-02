package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

type ConfigSQLite struct {
	DBPath      string
	Memory      bool
	CacheSize   int
	JournalMode string
}

func NewSQLiteDB(cfg *ConfigSQLite) (*sql.DB, error) {

	if cfg.DBPath == "" {
		return nil, fmt.Errorf("missing DBPath")
	}

	var dsn string
	if cfg.Memory {
		dsn = "file::memory:?cache=shard"
	} else {
		dbPath := strings.ReplaceAll(cfg.DBPath, "\\", "/")
		dsn = "file:" + dbPath
	}

	var params []string
	if cfg.CacheSize > 0 {
		params = append(params, fmt.Sprintf("_pragma=cache_size(%d)", cfg.CacheSize))
	}
	if cfg.JournalMode != "" {
		params = append(params, fmt.Sprintf("_pragma=journal_mode=%s", cfg.JournalMode))
	}

	if len(params) > 0 {
		dsn += "?" + joinDSNParams(params)
	}

	logrus.Debug("dsn: ", dsn)

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		logrus.Fatal("error open connection db:", err)
		return nil, err
	}

	logrus.Debug("db open")

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(30 * time.Minute)

	logrus.Debug("setting init")

	return db, nil
}

func joinDSNParams(params []string) string {
	return strings.Join(params, "&")
}
