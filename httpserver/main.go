package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// func NewInMemoryPlayerStore() *InMemoryPlayerStore {
// 	return &InMemoryPlayerStore{map[string]int{}}
// }

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	return s.scores[name], true
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.scores[name]++
}

func main() {
	// server := &PlayerServer{NewInMemoryPlayerStore()}
	// if err := http.ListenAndServe(":5000", server); err != nil {
	// 	log.Fatalf("Could not listen on port 5000, %v", err)
	// }

	//get login details
	var config Data
	GetConfig(&config)
	//Create the app
	app := &App{}
	app.ConnectAndPing(config)
	sqlStatement := `SELECT email from users where id = 1`
	rows, err := app.DB.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user.Email)
	}
}
