package db

import (
	"fmt"
	"os"
)

type Config struct {
	EnvKeyDBHost         string
	EnvKeyDBPort         string
	EnvKeyDBName         string
	EnvKeyDBUserName     string
	EnvKeyDBUserPassword string
}

// 外部から渡されてくる環境変数のキーの設定です
const (
	EnvKeyDBHost         = "DB_HOST"
	EnvKeyDBPort         = "DB_PORT"
	EnvKeyDBName         = "DB_NAME"
	EnvKeyDBUserName     = "DB_USER_NAME"
	EnvKeyDBUserPassword = "DB_USER_PASSWORD"
)

// 環境変数からデータベースの設定用の構造体を新規作成します
func NewConf() (c *Config, err error) {
	c = &Config{}

	c.EnvKeyDBHost = os.Getenv(EnvKeyDBHost)
	if c.EnvKeyDBHost == "" {
		return nil, fmt.Errorf("環境変数が設定されていません: %s", EnvKeyDBHost)
	}

	c.EnvKeyDBPort = os.Getenv(EnvKeyDBPort)
	if c.EnvKeyDBPort == "" {
		return nil, fmt.Errorf("環境変数が設定されていません: %s", EnvKeyDBPort)
	}

	c.EnvKeyDBName = os.Getenv(EnvKeyDBName)
	if c.EnvKeyDBName == "" {
		return nil, fmt.Errorf("環境変数が設定されていません: %s", EnvKeyDBName)
	}

	c.EnvKeyDBUserName = os.Getenv(EnvKeyDBUserName)
	if c.EnvKeyDBUserName == "" {
		return nil, fmt.Errorf("環境変数が設定されていません: %s", EnvKeyDBUserName)
	}

	c.EnvKeyDBUserPassword = os.Getenv(EnvKeyDBUserPassword)
	if c.EnvKeyDBUserPassword == "" {
		return nil, fmt.Errorf("環境変数が設定されていません: %s", EnvKeyDBUserPassword)
	}

	return
}
