package game_elements

import (
	//"math/rand"
	//"fmt"
	termbox "github.com/nsf/termbox-go"
)

type GameController struct {
	players PlayersQueue
	table Table
	direction bool
}

const STOP int = 10
const REVERSE int = 11
const PLUS_TWO int = 12
const PLUS_FOUR int = 13
const COLOR int = 14

func (gameController *GameController) actionOnTable(new_card Card) {
	new_color_on_board := new_card.color
	if new_card.action == PLUS_FOUR || new_card.action == COLOR {
		color_choosen := false
		new_color_stack := ColorStack{}
		new_color_stack.initialize()
		for !color_choosen {
			event := termbox.PollEvent()
			if event.Type == termbox.EventKey {
				if event.Key == termbox.KeyArrowLeft {
					new_color_stack.previousColor()
				} else if event.Key == termbox.KeyArrowRight {
					new_color_stack.nextColor()
				} else if event.Key == termbox.KeyEnter {
					new_color_on_board = new_color_stack.getColor()
					color_choosen = true
				}
			}

			termbox.Clear(termbox.ColorWhite, termbox.ColorWhite)
			gameController.players.renderBackward(70, 3)
			new_color_stack.render(30, 15)
			gameController.players.Top().renderHand(30, 30)
			gameController.table.ShowTable()
			termbox.Flush()
		}
	}
	if new_card.action == STOP {
		gameController.players.NextPlayer()
		gameController.players.NextPlayer()
	} else if new_card.action == REVERSE {
		gameController.players.reverseQueue()
	} else if new_card.action == PLUS_TWO {
		gameController.players.NextPlayer()
		gameController.players.Top().TakeCard(gameController.table)
		gameController.players.Top().TakeCard(gameController.table)
		gameController.players.NextPlayer()
	} else if new_card.action == PLUS_FOUR {
		gameController.players.NextPlayer()
		gameController.players.Top().TakeCard(gameController.table)
		gameController.players.Top().TakeCard(gameController.table)
		gameController.players.Top().TakeCard(gameController.table)
		gameController.players.Top().TakeCard(gameController.table)
		gameController.players.NextPlayer()
	} else {
		gameController.players.NextPlayer()
	}
	gameController.table.setNextAction(new_card.action)
	gameController.table.setNextColor(new_color_on_board)
}

func (gameController *GameController) StartGame() {
	gameController.direction = true
	gameController.table.Initialize()
	gameController.players.AddPlayer(Player{Name: "Bogdan"})
	gameController.players.AddPlayer(Player{Name: "Masha"})
	gameController.players.AddPlayer(Player{Name: "Anton"})
	gameController.players.AddPlayer(Player{Name: "Zamyatin"})
	for i := 0; i < gameController.players.GetSize(); i++ {
		for j := 0; j < 6; j++ {
			gameController.players.GetPlayer(i).TakeCard(gameController.table)
		}
	}
}

func (gameController *GameController) GameSession() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	end := false
	for !end {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {

			if event.Key == termbox.KeyEsc {
				end = true
			} else if event.Key == termbox.KeyArrowLeft {
				gameController.players.Top().previousCard()
			} else if event.Key == termbox.KeyArrowRight {
				gameController.players.Top().nextCard()
			} else if event.Key == termbox.KeyEnter {
				puted, new_card := gameController.players.Top().PutCard(&gameController.table)
				if puted {
					gameController.actionOnTable(new_card)
				}
			}
		}

		if !gameController.players.Top().checkAvailableCards(gameController.table) {
			gameController.players.Top().TakeCard(gameController.table)
			if !gameController.players.Top().checkAvailableCards(gameController.table) {
				gameController.players.NextPlayer()
			}
		}
		termbox.Clear(termbox.ColorWhite, termbox.ColorWhite)
		//fmt.Printf("Player %v\n", gameController.players.Top().Name)
		gameController.players.renderBackward(70, 3)
		gameController.players.Top().renderHand(30, 30)
		//fmt.Printf("Table\n")
		gameController.table.ShowTable()
		//fmt.Printf("---------------------player %v go------------------", gameController.players.Top().Name)
		//new_card := gameController.players.Top().PutCard(&gameController.table)
		//gameController.actionOnTable(new_card.action)
		termbox.Flush()
	}

}

func (gameController *GameController) IsGameEnd() bool{
	for i := 0; i < gameController.players.GetSize(); i++ {
		if gameController.players.GetPlayer(i).IsEmptyHand() {
			return true
		}
	}
	return false
}