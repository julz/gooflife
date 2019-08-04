package rules_test

import (
	"testing"

	"github.com/julz/gooflife/rules"
	"github.com/julz/gooflife/state"
	"gotest.tools/assert"
)

func TestApply(t *testing.T) {
	invert := rules.InvertRule{}

	next := invert.Apply(state.State{
		{false, true, true},
		{true, false, false},
		{false, true, true},
	})

	assert.DeepEqual(t, state.State{
		{true, false, false},
		{false, true, true},
		{true, false, false},
	}, next)
}
