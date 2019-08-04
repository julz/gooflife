package rules

import "github.com/julz/gooflife/state"

type RuleSet []CellRuleFunc

type CellRule interface {
	Apply(current state.CellState, neighbours int) state.CellState
}

func NewBasic() CellRule {
	return RuleSet{
		underpop,
		overpop,
		resurrect,
	}
}

func (rs RuleSet) Apply(current state.CellState, neighbours int) state.CellState {
	next := current
	for _, rule := range rs {
		next = rule(next, neighbours)
	}

	return next
}

type CellRuleFunc func(current state.CellState, neighbours int) state.CellState

func underpop(current state.CellState, neighbours int) state.CellState {
	if neighbours < 2 {
		return state.Dead
	}

	return current
}

func overpop(current state.CellState, neighbours int) state.CellState {
	if neighbours > 3 {
		return state.Dead
	}

	return current
}

func resurrect(current state.CellState, neighbours int) state.CellState {
	if neighbours == 3 {
		return state.Living
	}

	return current
}
