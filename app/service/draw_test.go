package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRow(t *testing.T) {
	m := [][]int{
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}

	l := [5]string{}

	for i, row := range m {
		l[i] = generateRow(row)
	}

	e := [5]string{
		`|#| | | | |`,
		`| |#| | | |`,
		`| | |#| | |`,
		`| | | |#| |`,
		`| | | | |#|`,
	}

	assert.Equal(t, l, e)
}
