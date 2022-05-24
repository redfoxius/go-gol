package main

import (
	"github.com/inancgumus/screen"
	"github.com/redfoxius/go-gol/app/models"
	"time"
)

func main() {
	screen.Clear()
	
	matrix := models.InitMatrix(25)
	matrix.Seed()
	matrix.Draw()

	for {
		screen.MoveTopLeft()

		time.Sleep(1 * time.Second)

		matrix.NextRound()

		if matrix.LifeIsDead() {
			break
		}
	}
}
