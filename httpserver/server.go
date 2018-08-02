package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := p.checkURL(w, req)
	if err != nil {
		return
	}
	constURL := "/players/"
	reqURL := req.URL.Path
	player := reqURL[len(constURL):]
	switch req.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.changeScore(w, player)
	}
}

func (p *PlayerServer) checkURL(w http.ResponseWriter, req *http.Request) error {
	constURL := "/players/"
	reqURL := req.URL.Path
	//Check len of request URL to make sure slice doesn't go out of range
	if len(reqURL) < len(constURL) {
		w.WriteHeader(http.StatusNotFound)
		return fmt.Errorf("url too short")
	}
	//Check if URL is valid starting with /players/
	if reqURL[0:len(constURL)] != constURL {
		w.WriteHeader(http.StatusNotFound)
		return fmt.Errorf("url is wrong, must start with /players/ ")
	}
	return nil
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, ok := p.store.GetPlayerScore(player)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) changeScore(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)

	w.WriteHeader(http.StatusAccepted)
}
