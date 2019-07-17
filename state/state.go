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
