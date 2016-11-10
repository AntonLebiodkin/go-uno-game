package game_elements

import "github.com/nsf/termbox-go"

type PlayersQueue struct {
	size  int
	Queue []Player
}

func (playersQueue *PlayersQueue) AddPlayer(player Player) {
	playersQueue.Queue = append(playersQueue.Queue, player)
	playersQueue.size = len(playersQueue.Queue)
}

func (playersQueue PlayersQueue) Top() *Player{
	return &playersQueue.Queue[0]
}

func (playersQueue *PlayersQueue) NextPlayer() *Player{
	last_player := playersQueue.Queue[0]
	for i := 1; i < len(playersQueue.Queue); i++ {
		playersQueue.Queue[i - 1] = playersQueue.Queue[i]
	}
	playersQueue.Queue[len(playersQueue.Queue) - 1] = last_player
	return &playersQueue.Queue[0];
}

func (playersQueue *PlayersQueue) ReverseQueue() {
	new_queue := []Player{}
	for i := len(playersQueue.Queue) - 1; i >= 0 ; i-- {
		new_queue = append(new_queue, playersQueue.Queue[i])
	}
	playersQueue.Queue = new_queue
}

func (playersQueue PlayersQueue) GetPlayer(index int) *Player{
	return &playersQueue.Queue[index]
}

func (playersQueue PlayersQueue) GetSize() int{
	return playersQueue.size
}

func (playersQueue PlayersQueue) renderBackward(x int, y int) {
	for i := 1; i < playersQueue.size; i++ {
		for j := 0; j < len(playersQueue.Queue[i].Name); j++ {
			termbox.SetCell(x + j, y + i * 5, rune(playersQueue.Queue[i].Name[j]), termbox.ColorBlack,
				termbox.ColorWhite)
		}
		playersQueue.Queue[i].renderBackwardHand(x + 10, y + i * 5)
	}
}