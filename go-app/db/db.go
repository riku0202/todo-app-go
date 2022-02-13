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
	var err error
	fmt.Printf("Openします")
	fmt.Printf(
		"%s:%s@tcp(%s:%s)/%s",
		c.EnvKeyDBUserName,
		c.EnvKeyDBUserPassword,
		c.EnvKeyDBHost,
		c.EnvKeyDBPort,
		c.EnvKeyDBName,
	)
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
		return nil, fmt.Errorf("\nOpen関数の呼び出しに失敗しました: %s\n", err)
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("\nPing:%s\n", err)
	}

	return db, nil
}
