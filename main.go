package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/venomspawn/sticks/ai"
	"github.com/venomspawn/sticks/game"
	"github.com/venomspawn/sticks/human"
)

func main() {
	fmt.Println(`Let's play a game of sticks! There are following rules:
1) a random amount of sticks (like |||||||) is generated at the very start;
2) you and me have turns one after one, taking one, two, or three sticks every
   time;
3) who takes the last sticks, loses the game;
4) who goes first is chosen randomly for every game;
5) you can input digits "1", "2", "3", or even sticks "|", "||", "|||"; all 
   whitespaces are ignored (so even input of " |    |     |    " is allowed!);
6) if your input is invalid, a corresponding error would be shown, after that
   you can try to input again;
`)

	rand.Seed(time.Now().UnixNano())

	sticks := 20 + rand.Intn(20)
	current := uint(rand.Intn(2))

	g, e := game.NewGame(sticks, 2, current)
	if e != nil {
		fmt.Printf("An error has appeared: %s\n", e)
		os.Exit(1)
	}

	s := strings.Repeat("|", g.Sticks())
	h := human.NewHuman()

	for !g.IsOver() {
		fmt.Println()
		fmt.Println("Current amount of sticks:")
		fmt.Println(s)

		if g.Current() == 0 {
			fmt.Println("My turn")
			e = g.TakeTurn(ai.Turn)
			if e == nil {
				t := len(s) - g.Sticks()
				fmt.Printf("I've taken %s\n", s[:t])
			}
		} else {
			fmt.Println("Your turn. How many sticks are you " +
				"gonna take?")
			e = g.TakeTurn(h.Turn)
		}

		if e != nil {
			fmt.Printf("An error has appeared: %s\n", e)
		}

		s = s[:g.Sticks()]
	}

	fmt.Println()
	if g.Current() == 0 {
		fmt.Println("I've won! ;-]")
	} else {
		fmt.Println("You've won! ;-(")
	}
}
