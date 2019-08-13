package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julz/gooflife/neighbours"
	"github.com/julz/gooflife/rules"
	"github.com/julz/gooflife/state"
)

func main() {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.Handle("/", http.FileServer(http.Dir(filepath.Join("ui", "build"))))

	http.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		s := state.State{}
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			panic(err) // let the http error handler catch this
		}

		neighbours := neighbours.WithWraparound(s)
		next := state.Apply(s, neighbours, rules.NewBasic())

		if err := json.NewEncoder(w).Encode(next); err != nil {
			panic(err) // let the http error handler catch this
		}
	})

	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
