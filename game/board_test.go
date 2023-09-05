package game

import (
	"fmt"
	"testing"
)

func TestBoard(t *testing.T) {
	board := NewBoard()
	board.SpawnBlock()
}

func TestShiftRight(t *testing.T) {
	cases := []struct {
		name   string
		start  *Board
		expect *Board
	}{
		{
			name:   "Shift Right empty board",
			start:  NewBoard(),
			expect: NewBoard(),
		},
		{
			name: "shift Right",
			start: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
				},
			},
			expect: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(0), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(0), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(0), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(0), NewTile(4)},
				},
			},
		},
		{
			name: "shift Right merge full",
			start: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(2), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(4)},
					{NewTile(2), NewTile(2), NewTile(2), NewTile(4)},
				},
			},
			expect: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(2), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(4)},
					{NewTile(0), NewTile(2), NewTile(4), NewTile(4)},
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			test.start.ShiftRight()

			for i := 0; i < len(test.start.tiles)-1; i++ {
				for j := 0; j < len(test.start.tiles[i])-1; j++ {
					if test.start.tiles[i][j].Value != test.expect.tiles[i][j].Value {
						fmt.Printf("%v", test.start)
						t.Errorf("start %v, expected %v", test.start.tiles[i][j], test.expect.tiles[i][j])
					}
				}
			}
		})
	}
}

func TestShiftLeft(t *testing.T) {
	cases := []struct {
		name    string
		start   *Board
		expect  *Board
		changed bool
	}{
		{
			name:    "Shift Left empty board",
			start:   NewBoard(),
			expect:  NewBoard(),
			changed: false,
		},
		{
			name: "Shift Left",
			start: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
				},
			},
			expect: &Board{
				tiles: [][]Tile{
					{NewTile(4), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(4), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(4), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(4), NewTile(0), NewTile(0), NewTile(0)},
				},
			},
			changed: true,
		},
		{
			name: "Shift Left No Action",
			start: &Board{
				tiles: [][]Tile{
					{NewTile(2), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(2), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(2), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(2), NewTile(0), NewTile(0), NewTile(0)},
				},
			},
			expect: &Board{
				tiles: [][]Tile{
					{NewTile(2), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(2), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(2), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(2), NewTile(0), NewTile(0), NewTile(0)},
				},
			},
			changed: false,
		},
		{
			name: "Shift Left full merge from left",
			start: &Board{
				tiles: [][]Tile{
					{NewTile(4), NewTile(2), NewTile(2), NewTile(2)},
					{NewTile(4), NewTile(2), NewTile(2), NewTile(2)},
					{NewTile(4), NewTile(2), NewTile(2), NewTile(2)},
					{NewTile(4), NewTile(2), NewTile(2), NewTile(2)},
				},
			},
			expect: &Board{
				tiles: [][]Tile{
					{NewTile(4), NewTile(4), NewTile(2), NewTile(0)},
					{NewTile(4), NewTile(4), NewTile(2), NewTile(0)},
					{NewTile(4), NewTile(4), NewTile(2), NewTile(0)},
					{NewTile(4), NewTile(4), NewTile(2), NewTile(0)},
				},
			},
			changed: true,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			changed := test.start.ShiftLeft()

			for i := 0; i < len(test.start.tiles)-1; i++ {
				for j := 0; j < len(test.start.tiles[i])-1; j++ {
					if test.start.tiles[i][j].Value != test.expect.tiles[i][j].Value {
						//fmt.Printf("%v", test.start)
						t.Errorf("start %v, expected %v", test.start.tiles[i][j], test.expect.tiles[i][j])
					}
				}
			}

			if changed != test.changed {
				t.Errorf("Flatten didn't return changed got:%v, wanted:%v", changed, test.changed)
			}
		})
	}
}

func TestShiftUp(t *testing.T) {
	cases := []struct {
		name   string
		start  *Board
		expect *Board
	}{
		{
			name:   "Shift Up empty board",
			start:  NewBoard(),
			expect: NewBoard(),
		},
		{
			name: "Shift Up",
			start: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
				},
			},
			expect: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(0), NewTile(0), NewTile(0), NewTile(0)},
				},
			},
		},
		{
			name: "Shift Up merge full",
			start: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
				},
			},
			expect: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(0), NewTile(0)},
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			test.start.ShiftUp()

			for i := 0; i < len(test.start.tiles)-1; i++ {
				for j := 0; j < len(test.start.tiles[i])-1; j++ {
					if test.start.tiles[i][j].Value != test.expect.tiles[i][j].Value {
						fmt.Printf("%v", test.start)
						t.Errorf("start %v, expected %v", test.start.tiles[i][j], test.expect.tiles[i][j])
					}
				}
			}
		})
	}
}

func TestShiftDown(t *testing.T) {
	cases := []struct {
		name   string
		start  *Board
		expect *Board
	}{
		{
			name:   "Shift Down empty board",
			start:  NewBoard(),
			expect: NewBoard(),
		},
		{
			name: "Shift Down",
			start: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
				},
			},
			expect: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(0), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
				},
			},
		},
		{
			name: "Shift Down full merge down",
			start: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
				},
			},
			expect: &Board{
				tiles: [][]Tile{
					{NewTile(0), NewTile(0), NewTile(0), NewTile(0)},
					{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
					{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
					{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			test.start.ShiftDown()
			for i := 0; i < len(test.start.tiles)-1; i++ {
				for j := 0; j < len(test.start.tiles[i])-1; j++ {
					if test.start.tiles[i][j].Value != test.expect.tiles[i][j].Value {
						t.Errorf("start %v, expected %v", test.start.tiles[i][j], test.expect.tiles[i][j])
					}
				}
			}
		})
	}
}
