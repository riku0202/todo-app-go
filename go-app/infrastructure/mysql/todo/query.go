package todo

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/riku0202/todo-app-go/go-app/infrastructure/mysql"
)

type Query struct {
	mysql.Infra
}

func NewQuery(DB *sql.DB) (*Query, error) {
	if DB == nil {
		return nil, fmt.Errorf("クエリを作成できません")
	}

	q := new(Query)
	q.DB = DB

	return q, nil
}

func (q *Query) FindByUserID(id string) (map[string]interface{}, error) {
	qs := fmt.Sprintf(
		`SELECT %s, %s, %s, %s FROM %s WHERE %s=?`,
		Conf.ColumnName.ID,
		Conf.ColumnName.Content,
		Conf.ColumnName.Created,
		Conf.ColumnName.Updated,
		Conf.TableName,
		Conf.ColumnName.UserID,
	)
	s, err := q.DB.Prepare(qs)
	if err != nil {
		return nil, fmt.Errorf("ステートメントの作成に失敗しました:%v", err)
	}
	defer q.CloseStmt(s)

	rows, err := s.Query(id)
	if err != nil {
		return nil, fmt.Errorf("クエリが失敗しました:%v", err)
	}

	var res []*Row
	for rows.Next() {
		ro := &Row{}
		err := rows.Scan(&ro.ID, &ro.Content, &ro.Created, &ro.Updated)
		if err != nil {
			return nil, fmt.Errorf("rowのデータを取得できません:%v", err)
		}
		res = append(res, ro)
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("rowのデータが存在しません:%v", err)
	}

	var data map[string]interface{}
	if err = json.Unmarshal([]byte(res[0].Content), &data); err != nil {
		return nil, fmt.Errorf("mapに変換できません:%v", err)
	}

	return data, nil
}
