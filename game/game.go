package game

import "errors"

type Game struct {
	sticks  int
	players uint
	current uint
}

var (
	ErrorInvalidPlayersAmount = errors.New("Players amount should be " +
		"a natural number, which is strictly more than one")

	ErrorInvalidCurrentPlayer = errors.New("Current player should be " +
		"non-negative number, which is strictly less than players " +
		"amount")
)

func NewGame(sticks int, players, current uint) (*Game, error) {
	if players == 0 {
		return nil, ErrorInvalidPlayersAmount
	}

	if players <= current {
		return nil, ErrorInvalidCurrentPlayer
	}

	return &Game{sticks: sticks, players: players, current: current}, nil
}

func (g *Game) Sticks() int {
	return g.sticks
}

func (g *Game) setSticks(sticks int) {
	g.sticks = sticks
}

func (g *Game) Players() uint {
	return g.players
}

func (g *Game) Current() uint {
	return g.current
}

func (g *Game) nextPlayer() {
	g.current = (g.Current() + 1) % g.Players()
}

func (g *Game) IsOver() bool {
	return g.Sticks() <= 0
}

const minSticksToTake = 1

func (g *Game) MinSticksToTake() int {
	return minSticksToTake
}

const maxSticksToTake = 3

func (g *Game) MaxSticksToTake() int {
	s := g.Sticks()

	if s <= 0 {
		return 0
	} else if s <= maxSticksToTake {
		return s
	} else {
		return maxSticksToTake
	}
}

var ErrorInvalidSticksToTake = errors.New("Invalid amount of sticks to take")

func (g *Game) TakeSticks(sticks int) error {
	if sticks < g.MinSticksToTake() || sticks > g.MaxSticksToTake() {
		return ErrorInvalidSticksToTake
	}

	g.setSticks(g.Sticks() - sticks)

	return nil
}

type Turn func(g *Game) error

var ErrorInvalidTurn = errors.New("Invalid turn: amount of sticks has not " +
	"been changed")

func (g *Game) TakeTurn(turn Turn) error {
	s := g.Sticks()

	e := turn(g)
	if e != nil {
		return e
	}

	if s == g.Sticks() {
		return ErrorInvalidTurn
	}

	g.nextPlayer()

	return nil
}
