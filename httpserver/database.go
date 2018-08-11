package main

import (
	"database/sql"
)

type App struct {
	DB *sql.DB
}

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

func (app *App) ConnectAndPing(config Data) error {
	var err error
	//Connect to database
	app.DB, err = sql.Open("postgres", config.psqlInfo)
	if err != nil {
		return err
	}
	err = app.DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

// func (app *App) QueryRetrieve(user *User, sqlStatement string) {
// 	row := app.DB.QueryRow(sqlStatement)
// 	row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
// }
