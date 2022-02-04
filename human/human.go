package human

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"

	"github.com/venomspawn/sticks/game"
)

type Human struct {
	r *bufio.Reader
}

func NewHuman() *Human {
	return &Human{r: bufio.NewReader(os.Stdin)}
}

var (
	whitespaces = regexp.MustCompile(`\s`)
	sticks = regexp.MustCompile(`\A\|+\z`)
	digit = regexp.MustCompile(`\A\d\z`)
)

var ErrorInvalidSticksInput = errors.New("Invalid amount of sticks to take")

func (h *Human) takeSticks() (int, error) {
	s, e := h.r.ReadString('\n')
	if e != nil && e != io.EOF {
		return 0, e
	}

	s = whitespaces.ReplaceAllString(s, "")
	switch {
	case digit.MatchString(s):
		return int(s[0]) - 48, nil
	case sticks.MatchString(s):
		return len(s), nil
	default:
		return 0, ErrorInvalidSticksInput
	}
}

var ErrorGameIsOver = errors.New("Can't make a turn: game is over")

func (h *Human) Turn(g *game.Game) error {
	if g.IsOver() {
		return ErrorGameIsOver
	}

	s, e := h.takeSticks()
	if e != nil {
		return e
	}

	return g.TakeSticks(s)
}
