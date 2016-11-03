package game_elements

import "github.com/nsf/termbox-go"

type Table struct {
	cardsStackOpened CardsStack
	cardsStackClosed CardsStack
	nextColor int
	nextAction int
}

func (table *Table) Initialize() {
	table.cardsStackClosed.Initialize()
	first_card := table.PopTopCard()
	for first_card.action > 9 {
		first_card = table.PopTopCard()
	}
	table.cardsStackOpened.PutCard(first_card)
	table.nextColor = first_card.color
	table.nextAction = first_card.action
}

func (table *Table) PopTopCard() Card{
	return table.cardsStackClosed.PopTopCard()
}

func (table Table) ShowTable() {
	next_card_color := termbox.ColorBlack
	if table.nextColor == 1 {
		next_card_color = termbox.ColorBlue
	} else if table.nextColor == 2 {
		next_card_color = termbox.ColorRed
	} else if table.nextColor == 3 {
		next_card_color = termbox.ColorYellow
	} else if table.nextColor == 4 {
		next_card_color = termbox.ColorGreen
	}
	if table.cardsStackOpened.TopCard().color != table.nextColor {
		termbox.SetCell(7, 7 , ' ', next_card_color, next_card_color)
		termbox.SetCell(8, 7 , ' ', next_card_color, next_card_color)
		termbox.SetCell(7, 8 , ' ', next_card_color, next_card_color)
		termbox.SetCell(8, 8 , ' ', next_card_color, next_card_color)
	}
	table.cardsStackOpened.TopCard().render(10, 10, false)
}

func (table *Table) PutCard(new_card Card) bool{
	if new_card.color == table.nextColor || new_card.color == 5 || new_card.action == table.nextAction || table.nextColor == 5{
		table.cardsStackOpened.PutCard(new_card)
		return true
	} else {
		return false
	} 
}

func (table *Table) setNextAction (newNextAction int) {
	table.nextAction = newNextAction
}

func (table *Table) setNextColor (newNextColor int) {
	table.nextColor = newNextColor
}