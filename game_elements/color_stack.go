package game_elements

type ColorStack struct {
	cards []Card
	choosen_color_index int
}

func (colorStack *ColorStack) initialize() {
	colorStack.choosen_color_index = 0
	colorStack.cards = append(colorStack.cards, Card{color: 1})
	colorStack.cards = append(colorStack.cards, Card{color: 2})
	colorStack.cards = append(colorStack.cards, Card{color: 3})
	colorStack.cards = append(colorStack.cards, Card{color: 4})
}

func (colorStack *ColorStack) nextColor() {
	colorStack.choosen_color_index++
	if colorStack.choosen_color_index == len(colorStack.cards) {
		colorStack.choosen_color_index = 0
	}
}

func (colorStack *ColorStack) previousColor() {
	colorStack.choosen_color_index--
	if colorStack.choosen_color_index == -1 {
		colorStack.choosen_color_index = len(colorStack.cards) - 1
	}
}

func (colorStack ColorStack) render(x int, y int) {
	for i := 0; i < len(colorStack.cards); i++ {
		if i == colorStack.choosen_color_index {
			colorStack.cards[i].render(i * 9 + x, y, true)
		} else {
			colorStack.cards[i].render(i * 9 + x, y, false)
		}
	}
}

func (colorStack ColorStack) getColor() int {
	return colorStack.cards[colorStack.choosen_color_index].color
}