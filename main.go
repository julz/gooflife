package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gosuri/uilive"
	"github.com/julz/gooflife/rules"
	"github.com/julz/gooflife/state"
)

func main() {
	s, err := state.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("parse input: %s", err)
	}

	game := rules.NewBasic()

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	for {
		s = state.Apply(s, state.Neighbours(s), game)
		fmt.Fprintln(writer, s)

		time.Sleep(1 * time.Second)
	}
}
