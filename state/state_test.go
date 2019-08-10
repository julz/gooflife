package state_test

import (
	"strings"
	"testing"

	"github.com/julz/gooflife/state"
	"gotest.tools/assert"
)

func TestApply(t *testing.T) {
	neighbours := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}

	next := state.Apply(state.State{
		{state.Living, state.Living, state.Living},
		{state.Living, state.Living, state.Living},
		{state.Living, state.Living, state.Living},
	}, neighbours, state.CellRuleFunc(func(c state.CellState, n int) state.CellState {
		if n == 3 {
			return state.Living
		}

		return state.Dead
	}))

	assert.DeepEqual(t, state.State{
		{state.Dead, state.Dead, state.Living},
		{state.Dead, state.Living, state.Dead},
		{state.Living, state.Dead, state.Dead},
	}, next)
}

func TestParse(t *testing.T) {
	table := []struct {
		Title    string
		Expected state.State
		Input    string
	}{
		{
			Title:    "single living cell",
			Expected: state.State{{state.Living}},
			Input:    "X",
		},
		{
			Title:    "single dead cell",
			Expected: state.State{{state.Dead}},
			Input:    ".",
		},
		{
			Title:    "various cells in a single row",
			Expected: state.State{{state.Dead, state.Living, state.Living}},
			Input:    ".XX",
		},
		{
			Title: "rows and columns",
			Expected: state.State{
				{state.Dead, state.Living, state.Living},
				{state.Living, state.Dead, state.Living},
			},
			Input: ".XX\nX.X",
		},
		{
			Title: "empty rows at ends",
			Expected: state.State{
				{state.Dead, state.Living, state.Living},
				{state.Living, state.Dead, state.Living},
			},
			Input: "\n.XX\nX.X\n",
		},
	}

	for _, eg := range table {
		example := eg
		t.Run(eg.Title, func(t *testing.T) {
			s, err := state.Parse(strings.NewReader(example.Input))
			assert.NilError(t, err)
			assert.DeepEqual(t, example.Expected, s)
		})
	}
}

func TestString(t *testing.T) {
	assert.Equal(t, ".", state.State{{state.Dead}}.String())
	assert.Equal(t, "X", state.State{{state.Living}}.String())
	assert.Equal(t, "X.X", state.State{{state.Living, state.Dead, state.Living}}.String())
	assert.Equal(t, ".XX\nX.X", state.State{
		{state.Dead, state.Living, state.Living},
		{state.Living, state.Dead, state.Living}}.String(),
	)
}
