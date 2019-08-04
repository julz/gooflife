package rules_test

import (
	"testing"

	"github.com/julz/gooflife/rules"
	"github.com/julz/gooflife/state"
	"gotest.tools/assert"
)

func TestBasicRules(t *testing.T) {
	table := []struct {
		title      string
		expect     state.CellState
		current    state.CellState
		neighbours int
	}{
		{
			title:      "underpopulation: 0 neighbours, should die",
			expect:     state.Dead,
			current:    state.Living,
			neighbours: 0,
		},
		{
			title:      "underpopulation: 1 neighbour, should die",
			expect:     state.Dead,
			current:    state.Living,
			neighbours: 1,
		},
		{
			title:      "underpopulation: 2 neighbours, should stay living",
			expect:     state.Living,
			current:    state.Living,
			neighbours: 2,
		},
		{
			title:      "survival: 2 neighbours, should stay dead",
			expect:     state.Dead,
			current:    state.Dead,
			neighbours: 2,
		},
		{
			title:      "survival: 2 neighbours, should stay living",
			expect:     state.Living,
			current:    state.Living,
			neighbours: 2,
		},
		{
			title:      "survival: 3 neighbours, should stay living",
			expect:     state.Living,
			current:    state.Living,
			neighbours: 3,
		},
		{
			title:      "resurrection: 3 neighbours, the dead should become living!",
			expect:     state.Living,
			current:    state.Dead,
			neighbours: 3,
		},
		{
			title:      "overpopulation: 4 neighbours, should die",
			expect:     state.Dead,
			current:    state.Living,
			neighbours: 4,
		},
		{
			title:      "overpopulation: 7 neighbours, should die",
			expect:     state.Dead,
			current:    state.Living,
			neighbours: 7,
		},
	}

	basic := rules.NewBasic()
	for _, eg := range table {
		ex := eg
		t.Run(eg.title, func(t *testing.T) {
			assert.Equal(t, ex.expect, basic.Apply(ex.current, ex.neighbours))
		})
	}
}
