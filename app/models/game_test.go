package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitMatrix(t *testing.T) {
	l := 5

	m := InitMatrix(l)

	assert.NotEmpty(t, m)
}

func TestGetEmptyMatrix(t *testing.T) {
	m1 := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	m2 := [][]int{
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}

	l := getEmptyMatrix(5)

	assert.NotEmpty(t, l)
	assert.Equal(t, l, m1)
	assert.NotEqual(t, l, m2)
}

func TestMatrix_Seed(t *testing.T) {
	m := InitMatrix(5)
	m.Seed()

	e := &matrix{mx: [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}}

	assert.Equal(t, m, e)
}

func TestMatrix_Check(t *testing.T) {
	l := 5

	m := InitMatrix(5)
	m.Seed()

	e := [][]int{
		{0, 1, 1, 1, 0},
		{0, 1, 1, 2, 1},
		{1, 3, 5, 3, 2},
		{1, 1, 3, 2, 2},
		{1, 2, 3, 2, 1},
	}

	n := getEmptyMatrix(l)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			n[i][j] = m.Check(i, j)
		}
	}

	assert.Equal(t, n, e)
}

func TestMatrix_NextRound(t *testing.T) {
	m := InitMatrix(5)
	m.Seed()
	m.NextRound()

	e := &matrix{mx: [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 1, 0, 0},
	}}

	assert.Equal(t, m, e)
}

func TestMatrix_LifeIsDead(t *testing.T) {
	m := InitMatrix(5)

	b1 := m.LifeIsDead()

	m.Seed()

	b2 := m.LifeIsDead()

	assert.Equal(t, b1, true)
	assert.Equal(t, b2, false)
	assert.NotEqual(t, b1, b2)
}
