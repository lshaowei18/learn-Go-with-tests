package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	return 123, true
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not listen on port 5000, %v", err)
	}
}