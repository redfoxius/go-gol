package models

import "github.com/redfoxius/go-gol/app/service"

type Matrix interface {
	Seed()
	Draw()
	NextRound()
	Check(x, y int) int
	LifeIsDead() bool
}

type matrix struct {
	mx [][]int
}

func InitMatrix(len int) Matrix {
	m := matrix{}
	m.mx = make([][]int, len)

	for i := range m.mx {
		m.mx[i] = make([]int, len)
	}

	return &m
}

func (m *matrix) Seed() {
	center := len(m.mx) / 2

	m.mx[center-1][center] = 1
	m.mx[center][center+1] = 1
	m.mx[center+1][center-1] = 1
	m.mx[center+1][center] = 1
	m.mx[center+1][center+1] = 1
}

func (m *matrix) Draw() {
	service.DrawMatrix(m.mx)
}

func (m *matrix) NextRound() {
	len := len(m.mx)

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			neighborsAlive := m.Check(i, j)

			switch {
			case m.mx[i][j] == 1 && neighborsAlive < 2:
				m.mx[i][j] = 0
			case m.mx[i][j] == 1 && neighborsAlive > 3:
				m.mx[i][j] = 0
			case m.mx[i][j] == 0 && neighborsAlive == 3:
				m.mx[i][j] = 1
			}
		}
	}

	m.Draw()
}

func (m *matrix) Check(x, y int) int {
	len := len(m.mx)
	sum := 0

	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i >= len {
			continue
		}
		for j := y - 1; j <= y+1; j++ {
			if j < 0 || j >= len {
				continue
			}

			sum += m.mx[i][j]
		}
	}

	return sum
}

func (m *matrix) LifeIsDead() bool {
	len := len(m.mx)

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if m.mx[i][j] == 1 {
				return false
			}
		}
	}

	return true
}
