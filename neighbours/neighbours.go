package neighbours

import (
	"github.com/julz/gooflife/state"
)

// WithoutWraparound returns the number of neighbours of each cell
// where neighbours that exceed the borders of the grid are ignored
func WithoutWraparound(s state.State) [][]int {
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
	result := make([][]int, len(s))
	for r, row := range s {
		result[r] = make([]int, len(row))
		for c := range row {

			n := 0
			for _, d := range deltas {
				if r+d.R < 0 {
					continue
				}
				if r+d.R >= len(s) {
					continue
				}
				if c+d.C < 0 {
					continue
				}
				if c+d.C >= len(row) {
					continue
				}

				if s[r+d.R][c+d.C] == state.Living {
					n++
				}
			}

			result[r][c] = n
		}
	}

	return result
}

// WithWraparound counts neighbours by wrapping around to the other
// side of the grid as needed
func WithWraparound(s state.State) [][]int {
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
	result := make([][]int, len(s))
	for r, row := range s {
		result[r] = make([]int, len(row))
		for c := range row {
			n := 0
			for _, d := range deltas {
				if s[wrap(r+d.R, len(s))][wrap(c+d.C, len(row))] {
					n++
				}
			}

			result[r][c] = n
		}
	}

	return result
}

func wrap(i, n int) int {
	if i < 0 {
		return n - 1
	}

	if i == n {
		return 0
	}

	return i
}
