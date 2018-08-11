package main

import (
	"testing"
)

//Test Initializing of database
func TestInitialize(t *testing.T) {
	app := App{} //new app struct
	t.Run("test bad config", func(t *testing.T) {

		//get login details
		var config Data

		err := app.ConnectAndPing(config)

		if err == nil {
			t.Errorf("Fail to test for empty psqlinfo")
		}
	})
	t.Run("correct config", func(t *testing.T) {

		//get login details
		var config Data
		GetConfig(&config)

		//Expects app.DB to be a connected DB
		err := app.ConnectAndPing(config)
		if err != nil {
			t.Errorf("%v", err)
		}

		if app.DB == nil {
			t.Errorf("Database not connected!")
		}
	})
}

// func TestQueryRetrieve(t *testing.T) {
// 	app := App{} //new app struct
// 	var config Data
// 	GetConfig(&config)
// 	app.ConnectAndPing(config)
// 	t.Run("Single Record", func(t *testing.T) {
// 		user := User{}
// 		sqlStatement := `SELECT * from users WHERE id=2;`
// 		app.QueryRetrieve(&user, sqlStatement)
// 		want := User{2, 22, "John", "Smith", "john@smith.com"}
// 		if !reflect.DeepEqual(user, want) {
// 			t.Errorf("got: %v, want: %v", user, want)
// 		}
// 	})

// t.Run("Multiple Records", func(t *testing.T) {
// 	want := User{}
// 	sqlStatement := `SELECT name FROM users LIMIT 2`
// 	app.QueryRetrieve(want, sqlStatement)
// 	got :=
// })
// }
