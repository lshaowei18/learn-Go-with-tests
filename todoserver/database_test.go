package main

import (
	"reflect"
	"testing"
)

type TodoData struct {
	ID       int
	Username string
	Todo     string
	Isdone   bool
}

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

func TestGetUserByName(t *testing.T) {
	t.Run("Get Pepper's Todo", func(t *testing.T) {
		app := App{}
		got := TodoData{}
		app.GetUserByName("Pepper", &got)
		want := TodoData{1, "Pepper", "Code", false}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
