package game

type Game struct {
	Board *Board
}

type Action string

func (g *Game) Start() {
	g.Board = NewBoard()
	g.Board.SpawnBlock()
}

func (g *Game) ProcessAction(move Action) {
	var changed bool = false
	switch string(move) {
	case "right":
		changed = g.Board.ShiftRight()
	case "left":
		changed = g.Board.ShiftLeft()
	case "down":
		changed = g.Board.ShiftDown()
	case "up":
		changed = g.Board.ShiftUp()
	}

	if changed {
		g.Board.SpawnBlock()
	}
}

func (g *Game) BoardString() string {
	return g.Board.String()
}
