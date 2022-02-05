package todo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/riku0202/todo-app-go/domain/model"
	"github.com/riku0202/todo-app-go/infrastructure/mysql"
)

type Repository struct {
	mysql.Infra
}

func NewRepository(DB *sql.DB) (*Repository, error) {
	if DB == nil {
		return nil, fmt.Errorf("値リポジトリを作成できません")
	}

	r := &Repository{}
	r.DB = DB

	return r, nil
}

func (r *Repository) Create(todo *model.Todo) error {
	qs := fmt.Sprintf(
		`INSERT INTO %s (%s, %s, %s, %s,%s) VALUES (?, ?, ?, ?, ?, ?)`,
		Conf.TableName,
		Conf.ColumnName.ID,
		Conf.ColumnName.UserID,
		Conf.ColumnName.Content,
		Conf.ColumnName.Created,
		Conf.ColumnName.Updated,
	)

	d, err := r.DB.Prepare(qs)
	if err != nil {
		return fmt.Errorf("ステートメントの作成に失敗しました")
	}
	defer func() { r.CloseStmt(d) }()

	j, err := json.Marshal(todo)
	if err != nil {
		return fmt.Errorf("TodoをJSONに変換できません")
	}

	_, err = d.Exec(
		todo.ID().String(),
		string(j),
		time.Now().Format("2022-01-01"),
		time.Now().Format("2022-01-01"),
	)
	if err != nil {
		return fmt.Errorf("クエリーが失敗しました:%v", err)
	}

	return nil
}

func (r *Repository) Delete(id *model.ID) error {
	qs := fmt.Sprintf(
		`DELETE FROM %s WHERE %s =?`,
		Conf.TableName,
		Conf.ColumnName.ID,
	)

	d, err := r.DB.Prepare(qs)
	if err != nil {
		return fmt.Errorf("ステートメントの作成に失敗しました:%v", err)
	}
	defer func() { r.CloseStmt(d) }()

	_, err = d.Exec(id.String())
	if err != nil {
		return fmt.Errorf("クエリーが失敗しました:%v", err)
	}

	return nil
}
