package main

import (
	"game_elements"
	"math/rand"
	"testing"
	"time"
	"strings"
)

func TestCardsStack(test *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	cardsStack := game_elements.CardsStack{}
	cardsStack.Initialize()
	if cardsStack.Size() == 108 {
		test.Log("Cards stack have 108 card.")
	} else {
		test.Errorf("Cards stack have %v card. 108 expected", cardsStack.Size())
	}
}


func TestPlayersOnBoard (test *testing.T) {
	playersQueue := game_elements.PlayersQueue{}
	playersQueue.AddPlayer(game_elements.Player{Name: "Bogdan"})
	playersQueue.AddPlayer(game_elements.Player{Name: "Masha"})
	playersQueue.AddPlayer(game_elements.Player{Name: "Anton"})
	playersQueue.AddPlayer(game_elements.Player{Name: "Zamyatin"})
	if playersQueue.GetSize() == 4 {
		test.Log("There are 4 players in the queue.")
	} else {
		test.Errorf("There is(are) %v player(s) in the queue. 4 expected", playersQueue.GetSize())
	}
}

func TestPlayerChanging(test *testing.T) {
	playersQueue := game_elements.PlayersQueue{}
	playersQueue.AddPlayer(game_elements.Player{Name: "Bogdan"})
	playersQueue.AddPlayer(game_elements.Player{Name: "Masha"})
	playersQueue.AddPlayer(game_elements.Player{Name: "Anton"})
	playersQueue.AddPlayer(game_elements.Player{Name: "Zamyatin"})
	current_player_name := playersQueue.Top().Name
	next_player_name := playersQueue.NextPlayer().Name
	if strings.Compare(next_player_name, current_player_name) == 0 {
		test.Error("Player changing at queue doesn't work")
	} else {
		test.Log("Players are changing in queue")
	}
}

func TestPlayersQueueReverse(test *testing.T) {
	playersQueue := game_elements.PlayersQueue{}
	playersQueue.AddPlayer(game_elements.Player{Name: "Bogdan"})
	playersQueue.AddPlayer(game_elements.Player{Name: "Masha"})
	playersQueue.AddPlayer(game_elements.Player{Name: "Anton"})
	playersQueue.AddPlayer(game_elements.Player{Name: "Zamyatin"})
	current_player_name := playersQueue.Queue[playersQueue.GetSize() - 1].Name
	playersQueue.ReverseQueue()
	next_player_name := playersQueue.Top().Name
	if strings.Compare(next_player_name, current_player_name) == 0 {
		test.Log("Queue reverse working well")
	} else {
		test.Error("Players reverse at queue doesn't work")
	}
}