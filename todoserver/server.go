package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type App struct {
	config Config
	db     *sql.DB
}

func UserServer(w http.ResponseWriter, r *http.Request) {
	//Handle URL to get user
	baseURL := r.URL.String()
	toRemove := "/users/"
	//Check URL to ensure that it won't slice bounds out of range
	if len(baseURL) < len(toRemove) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user := baseURL[len(toRemove):]

	switch user {
	case "Pepper":
		fmt.Fprintf(w, "Write Code")
	case "Shackleton":
		fmt.Fprintf(w, "Sail")
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
