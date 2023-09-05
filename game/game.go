package game

type Game struct {
	Board *Board
	State string
}

type Action string

func (g *Game) Start() {
	g.Board = NewBoard()
	g.State = "playing"
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

	if g.IsGameOver() {
		return
	}

	if changed {
		g.Board.SpawnBlock()
	}
}

func (g *Game) IsGameOver() bool {
	var emptyTiles bool = false

	for y := 0; y < len(g.Board.tiles); y++ {
		for x := 0; x < len(g.Board.tiles[y]); x++ {
			if g.Board.tiles[y][x] == 2048 {
				g.State = "win"
				return true
			}

			if g.Board.tiles[y][x] == 0 {
				emptyTiles = true
			}
		}
	}

	if !emptyTiles {
		g.State = "lose"
		return true
	}

	return false
}

func (g Game) BoardString() string {
	return g.Board.String()
}
