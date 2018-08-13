package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type TodoData struct {
	ID       int
	Username string
	Todo     string
	Isdone   bool
}

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

func (app App) GetTodoByUsername(name string) []TodoData {
	statement := `SELECT * FROM todolist WHERE username= $1`
	data := TodoData{}
	slice := []TodoData{}
	rows, err := app.db.Query(statement, name)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data.ID, &data.Username, &data.Todo, &data.Isdone)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows returned!")
		case nil:
			break
		default:
			log.Fatalln(err)
		}
		slice = append(slice, data)
	}
	return slice
}

func (app App) InsertTodo(username string, todo string) error {
	statement := `INSERT INTO todolist ( username, todo)
								VALUES($1, $2)`
	_, err := app.db.Exec(statement, username, todo)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
