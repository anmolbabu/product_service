package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
	"log"
	"sync"
)

var (
	dbClient    *sql.DB
	once        sync.Once
	cleanupFunc func() error
)

func GetPSQLClientPool(hostName string, port string, userName string, userPass string, dbName string) (*sql.DB, func() error, error) {
	var err error

	once.Do(func() {
		dbClient, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", hostName, port, userName, userPass, dbName))
		if err != nil {
			log.Fatal(err)
		}
		cleanupFunc = func() error {
			return dbClient.Close()
		}
	})

	return dbClient, cleanupFunc, err
}
