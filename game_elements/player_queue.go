package game_elements

import "github.com/nsf/termbox-go"

type PlayersQueue struct {
	size int
	queue []Player
}

func (playersQueue *PlayersQueue) AddPlayer(player Player) {
	playersQueue.queue = append(playersQueue.queue, player)
	playersQueue.size = len(playersQueue.queue)
}

func (playersQueue PlayersQueue) Top() *Player{
	return &playersQueue.queue[0]
}

func (playersQueue *PlayersQueue) NextPlayer() {
	last_player := playersQueue.queue[0]
	for i := 1; i < len(playersQueue.queue); i++ {
		playersQueue.queue[i - 1] = playersQueue.queue[i]
	}
	playersQueue.queue[len(playersQueue.queue) - 1] = last_player
}

func (playersQueue *PlayersQueue) reverseQueue() {
	new_queue := []Player{}
	for i := len(playersQueue.queue) - 1; i >= 0 ; i-- {
		new_queue = append(new_queue, playersQueue.queue[i])
	}
	playersQueue.queue = new_queue
}

func (playersQueue PlayersQueue) GetPlayer(index int) *Player{
	return &playersQueue.queue[index]
}

func (playersQueue PlayersQueue) GetSize() int{
	return playersQueue.size
}

func (playersQueue PlayersQueue) renderBackward(x int, y int) {
	for i := 1; i < playersQueue.size; i++ {
		for j := 0; j < len(playersQueue.queue[i].Name); j++ {
			termbox.SetCell(x + j, y + i * 5, rune(playersQueue.queue[i].Name[j]), termbox.ColorBlack,
				termbox.ColorWhite)
		}
		playersQueue.queue[i].renderBackwardHand(x + 10, y + i * 5)
	}
}