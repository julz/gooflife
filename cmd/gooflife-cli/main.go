package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/gosuri/uilive"
	"github.com/julz/gooflife/neighbours"
	"github.com/julz/gooflife/rules"
	"github.com/julz/gooflife/state"
)

func main() {
	wrap := flag.Bool("wrap", false, "count neighbours as if the grid wraps around")
	flag.Parse()

	neighbourFunc := neighbours.WithoutWraparound
	if *wrap {
		neighbourFunc = neighbours.WithWraparound
	}

	s, err := state.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("parse input: %s", err)
	}

	game := rules.NewBasic()

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	for {
		next := state.Apply(s, neighbourFunc(s), game)
		if reflect.DeepEqual(next, s) {
			return
		}

		s = next
		fmt.Fprintln(writer, s)
		time.Sleep(1 * time.Second)
	}
}
