package api

import "database/sql"

type ApplicationContext struct {
	client *sql.DB
}

func NewApplicationContext(client *sql.DB) ApplicationContext {
	return ApplicationContext{
		client: client,
	}
}
