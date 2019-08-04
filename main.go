package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/julz/gooflife/rules"
	"github.com/julz/gooflife/state"
)

func main() {
	s, err := state.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("parse input: %s", err)
	}

	game := rules.InvertRule{}

	for {
		s = game.Apply(s)
		fmt.Println(s)

		time.Sleep(1 * time.Second)
	}
}
