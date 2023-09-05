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

type Board struct {
	tiles [][]Tile
}

func NewBoard() *Board {
	var tiles [][]Tile = make([][]Tile, 4)
	for y := 0; y < 4; y++ {
		tiles[y] = make([]Tile, 4)
		for x := 0; x < 4; x++ {
			tiles[y][x] = NewTile(0)
		}
	}

	return &Board{
		tiles: tiles,
	}
}

func (b *Board) SpawnBlock() {
	for {
		y, x := rand.Intn(4), rand.Intn(4)

		if v := b.tiles[y][x]; v.empty {
			b.tiles[y][x] = NewTile(2)
			break
		}
	}
}

func (b *Board) ShiftRight() bool {
	var changedBoard bool
	for y := 0; y < 4; y++ {
		var changedRow bool
		b.tiles[y], changedRow = Flatten(b.tiles[y], "right")
		if changedRow {
			changedBoard = true
		}
	}

	return changedBoard
}

func (b *Board) ShiftLeft() bool {
	var changedBoard bool
	for y := 0; y < 4; y++ {
		var changedRow bool
		b.tiles[y], changedRow = Flatten(b.tiles[y], "left")
		if changedRow {
			changedBoard = true
		}
	}
	return changedBoard
}

func (b *Board) ShiftUp() bool {
	var changedBoard bool
	for x := 0; x < 4; x++ {
		var (
			tiles      []Tile = make([]Tile, 4)
			changedRow bool
		)
		for y := 0; y < 4; y++ {
			tiles[y] = b.tiles[y][x]
		}
		tiles, changedRow = Flatten(tiles, "left")
		if changedRow {
			changedBoard = true
		}

		for y := 0; y < len(b.tiles); y++ {
			b.tiles[y][x] = tiles[y]
		}
	}
	return changedBoard
}

func (b *Board) ShiftDown() bool {
	var changedBoard bool
	for x := 0; x < 4; x++ {
		var (
			tiles      []Tile = make([]Tile, 4)
			changedRow bool
		)
		for y := 0; y < 4; y++ {
			tiles[y] = b.tiles[y][x]
		}
		tiles, changedRow = Flatten(tiles, "right")
		if changedRow {
			changedBoard = true
		}
		for y := 0; y < len(b.tiles); y++ {
			b.tiles[y][x] = tiles[y]
		}
	}
	return changedBoard
}

func (b *Board) String() string {
	var output string

	for y := 0; y < len(b.tiles); y++ {
		for x := 0; x < 4; x++ {
			output += lipgloss.NewStyle().Foreground(colorMap[b.tiles[y][x].Value]).Render(fmt.Sprintf("%4d", b.tiles[y][x].Value))
		}
		output += "\n"
	}

	return output
}
