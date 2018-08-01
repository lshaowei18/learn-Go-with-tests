package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	constURL := "/players/"
	reqURL := req.URL.Path
	//Check len of request URL to make sure slice doesn't go out of range
	if len(reqURL) < len(constURL) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//Check if URL is valid starting with /players/
	if reqURL[0:len(constURL)] != constURL {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	player := reqURL[len(constURL):]
	score, ok := p.store.GetPlayerScore(player)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)
}
