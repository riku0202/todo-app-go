package mysql

import (
	"database/sql"
	"fmt"
)

type Infra struct {
	DB *sql.DB
}

func (i *Infra) CloseStmt(s *sql.Stmt) error {
	if s == nil {
		return fmt.Errorf("SQLステートメントを閉じることができません")
	}

	if err := s.Close(); err != nil {
		return fmt.Errorf("SQLステートメントを閉じることができません")
	}

	return nil
}
