package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/julz/gooflife/state"
)

func main() {
	s, err := state.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("parse input: %s", err)
	}

	game := InvertRule{}

	for {
		s = game.Apply(s)
		fmt.Println(s)

		time.Sleep(1 * time.Second)
	}
}

type InvertRule struct{}

func (InvertRule) Apply(previous state.State) state.State {
	result := make(state.State, len(previous))
	for r, row := range previous {
		result[r] = make([]state.CellState, len(row))
		for c, cell := range row {
			result[r][c] = !cell
		}
	}

	return result
}
