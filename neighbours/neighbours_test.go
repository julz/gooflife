package neighbours_test

import (
	"testing"

	"github.com/julz/gooflife/neighbours"
	"github.com/julz/gooflife/state"
	"gotest.tools/assert"
)

func TestWithoutWraparound(t *testing.T) {
	n := neighbours.WithoutWraparound(state.State{
		{state.Living, state.Living, state.Dead, state.Living, state.Living},
		{state.Dead, state.Living, state.Dead, state.Dead, state.Living},
		{state.Dead, state.Dead, state.Dead, state.Dead, state.Living},
		{state.Dead, state.Living, state.Living, state.Living, state.Dead},
	})

	assert.DeepEqual(t, n, [][]int{
		{2, 2, 3, 2, 2},
		{3, 2, 3, 4, 3},
		{2, 3, 4, 4, 2},
		{1, 1, 2, 2, 2},
	})
}

func TestWithWraparound(t *testing.T) {
	n := neighbours.WithWraparound(state.State{
		{state.Living, state.Living, state.Dead, state.Living},
		{state.Dead, state.Living, state.Dead, state.Living},
		{state.Dead, state.Dead, state.Dead, state.Living},
		{state.Dead, state.Living, state.Living, state.Dead},
	})

	assert.DeepEqual(t, n, [][]int{
		{5, 4, 6, 3},
		{6, 2, 5, 3},
		{4, 3, 5, 2},
		{5, 3, 4, 4},
	})
}
