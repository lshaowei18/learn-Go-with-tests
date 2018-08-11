package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func (app *App) ConnectDB() error {
	//Connect to database
	app.db, _ = sql.Open("postgres", app.config.sqlInfo)
	//Ping for database
	err := app.db.Ping()
	if err != nil {
		return err
	}
	return nil
}
