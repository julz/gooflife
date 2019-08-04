package state

import (
	"io"
	"io/ioutil"
	"strings"
)

type State [][]CellState

type CellState bool

const (
	Living CellState = true
	Dead   CellState = false
)

type CellRule interface {
	Apply(prev CellState, neighbours int) CellState
}

type CellRuleFunc func(current CellState, neighbours int) CellState

func (fn CellRuleFunc) Apply(c CellState, n int) CellState {
	return fn(c, n)
}

// Apply applies the rule to each cell in state, and returns the resulting new state
// the existing state is not modified
func Apply(state State, neighbours [][]int, rule CellRule) State {
	next := make(State, len(state))
	for r, row := range state {
		next[r] = make([]CellState, len(row))
		for c, cell := range row {
			next[r][c] = rule.Apply(cell, neighbours[r][c])
		}
	}

	return next
}

// Neighbours returns the number of neighbours of each cell
func Neighbours(state State) [][]int {
	deltas := []struct {
		R int
		C int
	}{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	result := make([][]int, len(state))
	for r, row := range state {
		result[r] = make([]int, len(row))
		for c := range row {

			n := 0
			for _, d := range deltas {
				if r+d.R < 0 {
					continue
				}
				if r+d.R >= len(row) {
					continue
				}
				if c+d.C < 0 {
					continue
				}
				if c+d.C >= len(row) {
					continue
				}

				if state[r+d.R][c+d.C] == Living {
					n++
				}
			}

			result[r][c] = n
		}
	}

	return result
}

// Parse parses an input stream in format "..X..XX\n.XX.XX" where . is a dead cell,
// and X is a live cell. Newlines separate rows.
func Parse(in io.Reader) (State, error) {
	s, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(s), "\n")

	state := make(State, len(lines))
	for l, line := range lines {
		state[l] = make([]CellState, 0)
		for _, c := range line {
			cell := Dead
			if c == 'X' {
				cell = Living
			}

			state[l] = append(state[l], cell)
		}
	}

	return state, nil
}

// String stringifies the state using the same format as Parse
func (s State) String() string {
	result := ""
	for _, line := range s {
		if result != "" {
			result += "\n"
		}

		for _, c := range line {
			if c == Living {
				result += "X"
			}

			if c == Dead {
				result += "."
			}
		}
	}

	return result
}
