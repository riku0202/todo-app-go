package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// CREATE TABLE todo_list (
// id VARCHAR(36) NOT NULL PRIMARY KEY,
// user_id VARCHAR(36) NOT NULL,
// title VARCHAR(200) NOT NULL,
// content VARCHAR(200) NOT NULL,
// finished BOOLEAN NOT NULL,
// created_at TIMESTAMP NOT NULL,
// updated_at TIMESTAMP NOT NULL
// );

type Todolist struct {
	Id        string
	UserId    string
	Title     string
	Content   string
	Finished  string
	CreatedAt string
	UpdatedAt string
}

func main() {
	fmt.Println("開始")
	db, err := initTCPConnectionPool()
	if err != nil {
		log.Fatalf("initTCPConnectionPool:%v", err)
	}
	fmt.Println("DB接続")

	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return
	} else {
		fmt.Println("データベース接続成功")
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	db, err := initTCPConnectionPool()
	if err != nil {
		log.Fatalf("initTCPConnectionPool:%v", err)
	}

	switch req.Method {
	case http.MethodGet:
		if err := getTodolist(w, req, db); err != nil {
			log.Fatalf("getTodoList:%v", err)
		}

	case http.MethodPost:
		if err := post(w, req, db); err != nil {
			log.Fatalf("postTodolist:%v", err)
		}

	case http.MethodDelete:
		if err := delete(w, req, db); err != nil {
			log.Fatalf("postTodolist:%v", err)
		}

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, fmt.Sprintf("HTTP Method %s Not Allowed", req.Method), http.StatusMethodNotAllowed)
	}
}

func post(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	if err := postTodolist(w, r, db); err != nil {
		log.Fatalf("postTodolist:%v", err)
	}
	if err := getTodolist(w, r, db); err != nil {
		log.Fatalf("getTodoList:%v", err)
	}

	return nil
}

func delete(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	if err := deleteTodolist(w, r, db); err != nil {
		log.Fatalf("deleteTodolist:%v", err)
	}
	if err := getTodolist(w, r, db); err != nil {
		log.Fatalf("getTodoList:%v", err)
	}

	return nil
}

func getTodolist(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	sqlSelect := "SELECT * FROM todo_list"

	sqlRows, err := db.Query(sqlSelect)
	if err != nil {
		return fmt.Errorf("DB.Query: %v", err)
	}

	defer sqlRows.Close()

	var rows []*Todolist

	for sqlRows.Next() {
		ro := &Todolist{}
		err := sqlRows.Scan(&ro.Id, &ro.Todo, &ro.Created, &ro.Updated)
		if err != nil {
			return fmt.Errorf("Rows.Scan: %v", err)
		}
		rows = append(rows, ro)
	}

	outjson, err := json.Marshal(rows)
	if err != nil {
		return fmt.Errorf("marshal: %v", err)
	}

	fmt.Fprint(w, string(outjson))

	return nil
}

func postTodolist(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	var jsonRow struct {
		Todo string
	}

	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)

	io.Copy(buf, body)

	if err := json.Unmarshal(buf.Bytes(), &jsonRow); err != nil {
		return fmt.Errorf("Unmarshal: %v", err)
	}

	sqlInsert := "INSERT INTO todo_list(id, todo, created, updated) VALUES(UUID(),?, NOW(), NOW())"

	if _, err := db.Exec(sqlInsert, jsonRow.Todo); err != nil {
		return fmt.Errorf("DB.Exec: %v", err)
	}

	return nil
}

func deleteTodolist(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	query := r.URL.Query()
	if query["id"] == nil {
		return nil
	}

	fmt.Println(query["id"][0])

	sqlDelete := "DELETE FROM todo_list WHERE id = ?"

	if _, err := db.Exec(sqlDelete, query["id"][0]); err != nil {
		return fmt.Errorf("DB.Exec: %v", err)
	}

	return nil
}

func initTCPConnectionPool() (*sql.DB, error) {
	var (
		dbUser    = mustGetenv("DB_USER") // e.g. 'my-db-user'
		dbPwd     = mustGetenv("DB_PASS") // e.g. 'my-db-password'
		dbTCPHost = mustGetenv("DB_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
		dbPort    = mustGetenv("DB_PORT") // e.g. '3306'
		dbName    = mustGetenv("DB_NAME") // e.g. 'my-database'
	)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPwd, dbTCPHost, dbPort, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	configureConnectionPool(dbPool)

	return dbPool, nil
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func configureConnectionPool(dbPool *sql.DB) {
	dbPool.SetMaxIdleConns(5)

	dbPool.SetMaxOpenConns(7)

	dbPool.SetConnMaxLifetime(1800 * time.Second)
}
