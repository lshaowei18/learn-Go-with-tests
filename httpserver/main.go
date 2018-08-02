package main

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
	// server := &PlayerServer{&InMemoryPlayerStore{}}
	// if err := http.ListenAndServe(":5000", server); err != nil {
	// 	log.Fatalf("Could not listen on port 5000, %v", err)
	// }
}
