// package db

// import "os"

// func init() (*sql.DB, error) {
// 	var (
// 		dbUser    = mustGetenv("DB_USER") // e.g. 'my-db-user'
// 		dbPwd     = mustGetenv("DB_PASS") // e.g. 'my-db-password'
// 		dbTCPHost = mustGetenv("DB_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
// 		dbPort    = mustGetenv("DB_PORT") // e.g. '3306'
// 		dbName    = mustGetenv("DB_NAME") // e.g. 'my-database'
// 	)

// 	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPwd, dbTCPHost, dbPort, dbName)

// 	dbPool, err := sql.Open("mysql", dbURI)
// 	if err != nil {
// 		return nil, fmt.Errorf("sql.Open: %v", err)
// 	}

// 	configureConnectionPool(dbPool)

// 	return dbPool, nil
// }

// func configureConnectionPool(dbPool *sql.DB) {
// 	dbPool.SetMaxIdleConns(5)

// 	dbPool.SetMaxOpenConns(7)

// 	dbPool.SetConnMaxLifetime(1800 * time.Second)
// }

// func mustGetenv(k string) string {
// 	v := os.Getenv(k)
// 	if v == "" {
// 		log.Fatalf("Warning: %s environment variable not set.\n", k)
// 	}
// 	return v
// }
