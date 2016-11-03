package game_elements

import (
	"fmt"
	//"strconv"
	//"github.com/fatih/color"
	termbox "github.com/nsf/termbox-go"
)

type Card struct{
	color int
	action int
}

func (card Card) Print(){
	fmt.Printf("Color: %v Action: %v\n", card.color, card.action)
}

func (card Card) render(x int, y int, choosen bool){
	card_fg_color := termbox.ColorBlack
	if card.color == 1 {
		card_fg_color = termbox.ColorBlue
	} else if card.color == 2 {
		card_fg_color = termbox.ColorRed
	} else if card.color == 3 {
		card_fg_color = termbox.ColorYellow
	} else if card.color == 4 {
		card_fg_color = termbox.ColorGreen
	} else if card.color == 5 {
		card_fg_color = termbox.ColorBlack
	}
	if choosen {
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				termbox.SetCell(x + j - 1, y + i - 1, ' ', card_fg_color, card_fg_color)
			}
		}
	}
	if card.action == 10 {
		for i := 0; i < 6; i++ {
			termbox.SetCell(x + i, y, ' ', card_fg_color, card_fg_color)
		}

		termbox.SetCell(x, y + 1, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 1, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 2, y + 1, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 3, y + 1, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 4, y + 1, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 5, y + 1, ' ', card_fg_color, card_fg_color)

		termbox.SetCell(x, y + 2, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 2, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 2, y + 2, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 3, y + 2, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 4, y + 2, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 5, y + 2, ' ', card_fg_color, card_fg_color)

		termbox.SetCell(x, y + 3, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 3, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 2, y + 3, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 3, y + 3, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 4, y + 3, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 5, y + 3, ' ', card_fg_color, card_fg_color)

		termbox.SetCell(x, y + 4, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 4, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 2, y + 4, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 3, y + 4, ' ', termbox.ColorBlack, termbox.ColorBlack)
		termbox.SetCell(x + 4, y + 4, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 5, y + 4, ' ', card_fg_color, card_fg_color)

		for i := 0; i < 6; i++ {
			termbox.SetCell(x + i, y + 5, ' ', card_fg_color, card_fg_color)
		}
	} else if card.action == 11 {
		for i := 0; i < 6; i++ {
			termbox.SetCell(x + i, y, ' ', card_fg_color, card_fg_color)
		}

		termbox.SetCell(x, y + 1, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 1, '\u256d', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 2, y + 1, '\u2501', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 3, y + 1, '\u2501', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 4, y + 1, '\u2191', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 5, y + 1, ' ', card_fg_color, card_fg_color)

		termbox.SetCell(x, y + 2, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 2, '\u2502', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 2, y + 2, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 3, y + 2, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 4, y + 2, '\u2502', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 5, y + 2, ' ', card_fg_color, card_fg_color)

		termbox.SetCell(x, y + 3, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 3, '\u2502', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 2, y + 3, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 3, y + 3, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 4, y + 3, '\u2502', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 5, y + 3, ' ', card_fg_color, card_fg_color)

		termbox.SetCell(x, y + 4, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 4, '\u2193', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 2, y + 4, '\u2501', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 3, y + 4, '\u2501', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 4, y + 4, '\u256f', termbox.ColorBlack, card_fg_color)
		termbox.SetCell(x + 5, y + 4, ' ', card_fg_color, card_fg_color)

		for i := 0; i < 6; i++ {
			termbox.SetCell(x + i, y + 5, ' ', card_fg_color, card_fg_color)
		}
	} else if card.action == 12 || card.action == 13{
		for i := 0; i < 6; i++ {
			for j := 0; j < 6; j++ {
				if i == 2 && j == 2 {
					if card.action == 12 {
						termbox.SetCell(x + i, y + j, '+', termbox.ColorBlack, card_fg_color)
					} else if card.action == 13 {
						termbox.SetCell(x + i, y + j, '+', termbox.ColorWhite, card_fg_color)
					}
				} else if i == 3 && j == 2 {
					if card.action == 12 {
						termbox.SetCell(x + i, y + j, '2', termbox.ColorBlack, card_fg_color)
					} else if card.action == 13 {
						termbox.SetCell(x + i, y + j, '4', termbox.ColorWhite, card_fg_color)
					}
				} else {
					termbox.SetCell(x + i, y + j, ' ', card_fg_color, card_fg_color)
				}
			}
		}
	} else if card.action == 14 {
		for i := 0; i < 6; i++ {
			termbox.SetCell(x + i, y, ' ', card_fg_color, card_fg_color)
		}

		termbox.SetCell(x, y + 1, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 1, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 2, y + 1, ' ', termbox.ColorBlue, termbox.ColorBlue)
		termbox.SetCell(x + 3, y + 1, ' ', termbox.ColorGreen, termbox.ColorGreen)
		termbox.SetCell(x + 4, y + 1, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 5, y + 1, ' ', card_fg_color, card_fg_color)

		termbox.SetCell(x, y + 2, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 2, ' ', termbox.ColorBlue, termbox.ColorBlue)
		termbox.SetCell(x + 2, y + 2, ' ', termbox.ColorBlue, termbox.ColorBlue)
		termbox.SetCell(x + 3, y + 2, ' ', termbox.ColorGreen, termbox.ColorGreen)
		termbox.SetCell(x + 4, y + 2, ' ', termbox.ColorGreen, termbox.ColorGreen)
		termbox.SetCell(x + 5, y + 2, ' ', card_fg_color, card_fg_color)

		termbox.SetCell(x, y + 3, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 3, ' ', termbox.ColorRed, termbox.ColorRed)
		termbox.SetCell(x + 2, y + 3, ' ', termbox.ColorRed, termbox.ColorRed)
		termbox.SetCell(x + 3, y + 3, ' ', termbox.ColorYellow, termbox.ColorYellow)
		termbox.SetCell(x + 4, y + 3, ' ', termbox.ColorYellow, termbox.ColorYellow)
		termbox.SetCell(x + 5, y + 3, ' ', card_fg_color, card_fg_color)

		termbox.SetCell(x, y + 4, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 1, y + 4, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 2, y + 4, ' ', termbox.ColorRed, termbox.ColorRed)
		termbox.SetCell(x + 3, y + 4, ' ', termbox.ColorYellow, termbox.ColorYellow)
		termbox.SetCell(x + 4, y + 4, ' ', card_fg_color, card_fg_color)
		termbox.SetCell(x + 5, y + 4, ' ', card_fg_color, card_fg_color)

		for i := 0; i < 6; i++ {
			termbox.SetCell(x + i, y + 5, ' ', card_fg_color, card_fg_color)
		}
	} else {
		for i := 0; i < 6; i++ {
			for j := 0; j < 6; j++ {
				if i == 3 && j == 2 {
					switch card.action {
					case 0:
						termbox.SetCell(x + i, y + j, '0', termbox.ColorBlack, card_fg_color)
					case 1:
						termbox.SetCell(x + i, y + j, '1', termbox.ColorBlack, card_fg_color)
					case 2:
						termbox.SetCell(x + i, y + j, '2', termbox.ColorBlack, card_fg_color)
					case 3:
						termbox.SetCell(x + i, y + j, '3', termbox.ColorBlack, card_fg_color)
					case 4:
						termbox.SetCell(x + i, y + j, '4', termbox.ColorBlack, card_fg_color)
					case 5:
						termbox.SetCell(x + i, y + j, '5', termbox.ColorBlack, card_fg_color)
					case 6:
						termbox.SetCell(x + i, y + j, '6', termbox.ColorBlack, card_fg_color)
					case 7:
						termbox.SetCell(x + i, y + j, '7', termbox.ColorBlack, card_fg_color)
					case 8:
						termbox.SetCell(x + i, y + j, '8', termbox.ColorBlack, card_fg_color)
					case 9:
						termbox.SetCell(x + i, y + j, '9', termbox.ColorBlack, card_fg_color)
					}
				} else {
					termbox.SetCell(x + i, y + j, ' ', card_fg_color, card_fg_color)
				}
			}
		}
	}
}

func (card Card) renderBackward(x int, y int){
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			termbox.SetCell(x + j, y + i, ' ', termbox.ColorMagenta, termbox.ColorMagenta)
		}
	}
}

func (card Card) GetColor() int{
	return card.color
}

func (card Card) GetAction() int{
	return card.action
}