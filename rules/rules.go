package rules

import "github.com/julz/gooflife/state"

type RuleSet []Rule

type Rule interface {
	Apply(previous state.State) state.State
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
