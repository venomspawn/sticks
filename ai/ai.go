package ai

import (
	"errors"
	"math/rand"

	"github.com/venomspawn/sticks/game"
)

var ErrorGameIsOver = errors.New("Can't make a turn: game is over")

func Turn(g *game.Game) error {
	s := g.Sticks()

	if s <= 0 {
		return ErrorGameIsOver
	}

	r := (s - 1) & 3
	if r == 0 {
		r = g.MinSticksToTake() + rand.Intn(g.MaxSticksToTake())
	}

	return g.TakeSticks(r)
}
