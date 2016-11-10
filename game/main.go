package main

import (
	"game_elements"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	game := game_elements.GameController{}
	game.StartGame()
	game.GameSession()
}
