package main

import (
	"reflect"
	"testing"
)

func TestConnectDB(t *testing.T) {
	t.Run("Without Config", func(t *testing.T) {
		app := App{}
		err := app.ConnectDB()

		if err == nil {
			t.Errorf("Should fail, no user config details.")
		}
	})
	t.Run("With Config", func(t *testing.T) {
		app := App{}
		app.GetConfig()
		app.ConnectDB()

		if app.db == nil {
			t.Errorf("ConnectDB failed.")
		}
	})
}

func TestGetTodoByUsername(t *testing.T) {
	app := App{}
	app.GetConfig()
	app.ConnectDB()
	t.Run("Get Pepper's Todo", func(t *testing.T) {
		got := app.GetTodoByUsername("Pepper")
		want := []TodoData{{1, "Pepper", "Code", false}}
		checkSlice(t, got, want)
	})
	t.Run("Get multiple todos", func(t *testing.T) {
		got := app.GetTodoByUsername("Shackleton")
		want := []TodoData{{6, "Shackleton", "Sail", false}, {7, "Shackleton", "Remain positive", false}}
		checkSlice(t, got, want)
	})
}

func TestInsertTodo(t *testing.T) {
	app := App{}
	app.GetConfig()
	app.ConnectDB()
	t.Run("Insert Carson todo", func(t *testing.T) {
		err := app.InsertTodo("Carson", "Research on sea")
		if err != nil {
			t.Errorf("Insert failed.")
		}
	})
}

func TestDeleteTodo(t *testing.T) {
	app := App{}
	app.GetConfig()
	app.ConnectDB()

	t.Run("Delete Carson todo", func(t *testing.T) {
		app.InsertTodo("Carson", "")
	})
}

func checkSlice(t *testing.T, got []TodoData, want []TodoData) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
