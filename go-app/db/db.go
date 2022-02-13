package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mu   sync.Mutex
	db   *sql.DB
	open = func(driverName, dataSourceName string) (*sql.DB, error) {
		return sql.Open(driverName, dataSourceName)
	}
)

func NewDB(c *Config) (*sql.DB, error) {
	mu.Lock()
	defer mu.Unlock()

	if db != nil {
		return db, nil
	}
	fmt.Printf("Openします")
	var err error
	db, err = open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			c.EnvKeyDBUserName,
			c.EnvKeyDBUserPassword,
			c.EnvKeyDBHost,
			c.EnvKeyDBPort,
			c.EnvKeyDBName,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("Open関数の呼び出しに失敗しましたA: %s", err)
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Ping関数の呼び出しに失敗しました: %s", err)
	}

	return db, nil
}
