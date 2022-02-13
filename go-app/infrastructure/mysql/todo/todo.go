package todo

import "fmt"

type Config struct {
	TableName  string
	ColumnName *ColumnName
}

// テーブルのカラム名です
type ColumnName struct {
	ID      string
	UserID  string
	Content string
	Created string
	Updated string
}

// 結果の行です
type Row struct {
	ID      string
	UserID  string
	Content string
	Created string
	Updated string
}

var c = &Config{
	TableName: "billings",
	ColumnName: &ColumnName{
		ID:      "id",
		UserID:  "user_id",
		Content: "content",
		Created: "created",
		Updated: "updated",
	},
}

// テーブルの設定を返します
func GetConfig() *Config {
	return c
}

var Conf = GetConfig()

var NotFoundError = fmt.Errorf("TODOが見つかりません")
