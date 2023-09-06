package game

import (
	"fmt"
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

var (
	colorMap map[int]lipgloss.Color = map[int]lipgloss.Color{
		0:    lipgloss.Color("#ffffff"),
		2:    lipgloss.Color("#FF4500"),
		4:    lipgloss.Color("#FFA500"),
		8:    lipgloss.Color("#FFD700"),
		16:   lipgloss.Color("#FFFF00"),
		32:   lipgloss.Color("#00FF00"),
		64:   lipgloss.Color("#00FFFF"),
		128:  lipgloss.Color("#0000FF"),
		256:  lipgloss.Color("#4B0082"),
		512:  lipgloss.Color("#8A2BE2"),
		1080: lipgloss.Color("#9400D3"),
		2040: lipgloss.Color("#8B0000"),
	}
)

type Board []TileRow

func NewBoard() Board {
	var board Board = make(Board, 4)
	for y := 0; y < 4; y++ {
		board[y] = make(TileRow, 4)
		for x := 0; x < 4; x++ {
			board[y][x] = Tile(0)
		}
	}

	return board
}

func (b Board) SpawnBlock() {
	for {
		y, x := rand.Intn(4), rand.Intn(4)

		if v := b[y][x]; v == 0 {
			b[y][x] = Tile(2)
			break
		}
	}
}

func (b Board) ShiftRight() bool {
	var changedBoard bool
	for y := 0; y < 4; y++ {
		var changedRow bool
		b[y], changedRow = Flatten(b[y], "right")
		if changedRow {
			changedBoard = true
		}
	}

	return changedBoard
}

func (b Board) ShiftLeft() bool {
	var changedBoard bool
	for y := 0; y < 4; y++ {
		var changedRow bool
		b[y], changedRow = Flatten(b[y], "left")
		if changedRow {
			changedBoard = true
		}
	}
	return changedBoard
}

func (b Board) ShiftUp() bool {
	var changedBoard bool
	for x := 0; x < 4; x++ {
		var (
			tileRow    TileRow = b.getColumn(x)
			changedRow bool
		)
		tileRow, changedRow = Flatten(tileRow, "left")
		if changedRow {
			changedBoard = true
		}

		for y := 0; y < len(b); y++ {
			b[y][x] = tileRow[y]
		}
	}
	return changedBoard
}

func (b Board) ShiftDown() bool {
	var changedBoard bool
	for x := 0; x < 4; x++ {
		var (
			tileRow    TileRow = b.getColumn(x)
			changedRow bool
		)
		tileRow, changedRow = Flatten(tileRow, "right")
		if changedRow {
			changedBoard = true
		}
		for y := 0; y < len(b); y++ {
			b[y][x] = tileRow[y]
		}
	}
	return changedBoard
}

func (b Board) getColumn(col int) TileRow {
	var tileColumn TileRow
	for y := 0; y < 4; y++ {
		tileColumn = append(tileColumn, b[y][col])
	}
	return tileColumn
}

func (b Board) String() string {
	var output string

	for y := 0; y < len(b); y++ {
		for x := 0; x < 4; x++ {
			output += lipgloss.NewStyle().Foreground(colorMap[int(b[y][x])]).Render(fmt.Sprintf("%5d", b[y][x]))
		}
		output += "\n"
	}

	return output
}
