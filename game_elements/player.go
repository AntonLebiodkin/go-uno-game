package game_elements

import (
	"fmt"
)

type Player struct{
	Name string
	cards []Card
	choosen_card_index int
}

func (player *Player) PutCard(table *Table) (bool, Card) {
	if table.PutCard(player.cards[player.choosen_card_index]) {
		putted_card := player.cards[player.choosen_card_index]
		for i := player.choosen_card_index; i < len(player.cards) - player.choosen_card_index - 1; i++ {
			player.cards[i] = player.cards[i + 1]
		}
		player.cards = player.cards[:len(player.cards) - 1]
		player.choosen_card_index = 0
		return true, putted_card
	}
	return false, Card{}
}

func (player *Player) nextCard() {
	player.choosen_card_index++
	if player.choosen_card_index == len(player.cards) {
		player.choosen_card_index = 0
	}
}

func (player *Player) previousCard() {
	player.choosen_card_index--
	if player.choosen_card_index == -1 {
		player.choosen_card_index = len(player.cards) - 1
	}
}

func (player Player) renderHand(x int, y int) {
	for i := 0; i < len(player.cards); i++ {
		if i == player.choosen_card_index {
			player.cards[i].render(i * 9 + x, y, true)
		} else {
			player.cards[i].render(i * 9 + x, y, false)
		}
	}
}

func (player Player) checkAvailableCards(table Table) bool{
	for i := 0; i < len(player.cards); i++ {
		if player.cards[i].color == 5 || player.cards[i].color == table.nextColor || player.cards[i].action == table.nextAction{
			return true
		}
	}
	return false
}

func (player Player) renderBackwardHand(x int, y int) {
	for i := 0; i < len(player.cards); i++ {
		player.cards[i].renderBackward(i * 5 + x, y)
	}
}

func (player *Player) TakeCard(table Table) {
	new_card := table.PopTopCard()
	player.cards = append(player.cards, new_card)
}

func (player Player) PrintCards() {
	for i := 0; i < len(player.cards); i++ {
		player.cards[i].Print()
	}
	fmt.Println()
}

func (player Player) IsEmptyHand() bool{
	if len(player.cards) == 0 {
		return true
	} else {
		return false
	}
}