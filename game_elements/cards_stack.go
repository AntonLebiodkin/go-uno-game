package game_elements

import (
	"math/rand"
)

type CardsStack struct {
	cards []Card
}

func (cardsStack *CardsStack) PutCard(new_card Card) {
	cardsStack.cards = append(cardsStack.cards, Card{})
	for i := len(cardsStack.cards) - 1; i > 0; i-- {
		cardsStack.cards[i] = cardsStack.cards[i - 1]
	}
	cardsStack.cards[0] = new_card
}

func (cardsStack *CardsStack) Print() {
	for i := 0; i < len(cardsStack.cards); i++ {
		cardsStack.cards[i].Print()
	}
}

func (cardsStack *CardsStack) Initialize() {
	for i := 0; i < 4; i++ {
	    cardsStack.cards = append(cardsStack.cards, Card{action: 0, color: i + 1})
	    cardsStack.cards = append(cardsStack.cards, Card{action: 13, color: 5})
	    cardsStack.cards = append(cardsStack.cards, Card{action: 14, color: 5})
	}
	for i := 0; i < 12; i++ {
		for j := 0; j < 4; j++ {
			cardsStack.cards = append(cardsStack.cards, Card{action: i + 1, color: j + 1})
			cardsStack.cards = append(cardsStack.cards, Card{action: i + 1, color: j + 1})
		}
	}
	shuffled_cards := []Card{}
	perm := rand.Perm(len(cardsStack.cards))
	for _, v := range perm {
	    shuffled_cards = append(shuffled_cards, cardsStack.cards[v])
	}
	cardsStack.cards = shuffled_cards
}

func (cardsStack *CardsStack) PopTopCard() Card{
	if len(cardsStack.cards) != 0 {
		top_card := cardsStack.cards[0]
		size := len(cardsStack.cards)
		for i := 0; i < size - 1; i++ {
			cardsStack.cards[i] = cardsStack.cards[i + 1]
		}
		cardsStack.cards = cardsStack.cards[:size - 1]
		return top_card
	}
	return Card{}
}

func (cardsStack CardsStack) TopCard() Card {
	if len(cardsStack.cards) != 0 {
		return cardsStack.cards[0]
	}
	return Card{}	
}

func (cardsStack CardsStack) Size() int {
	return len(cardsStack.cards)
}