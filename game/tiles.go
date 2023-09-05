package game

type Tile struct {
	Value int
	empty bool
}

func NewTile(v int) Tile {
	var empty bool = true
	if v != 0 {
		empty = false
	}
	return Tile{
		Value: v,
		empty: empty,
	}
}

func (t *Tile) Double() {
	t.Value = t.Value * 2
}

func (t *Tile) Empty() {
	t.Value = 0
	t.empty = true
}

func Flatten(tiles []Tile, direction string) ([]Tile, bool) {
	var changed bool = false
	flatStack := RemoveEmpty(tiles)

	switch direction {
	case "right":
		for i := len(flatStack) - 1; i > 0; i-- {
			if flatStack[i].Value == flatStack[i-1].Value {
				flatStack[i].Double()
				flatStack[i-1] = NewTile(0)
			}
		}
	case "left":
		for i := 0; len(flatStack)-1 > i; i++ {
			if flatStack[i].Value == flatStack[i+1].Value {
				flatStack[i].Double()
				flatStack[i+1] = NewTile(0)
			}
		}
	}

	flatStack = RemoveEmpty(flatStack)

	for len(flatStack) < 4 {
		if direction == "right" {
			flatStack = append([]Tile{NewTile(0)}, flatStack...)
		} else {
			flatStack = append(flatStack, NewTile(0))
		}
	}

	for i := 0; i < 4; i++ {
		if tiles[i] != flatStack[i] {
			changed = true
			break
		}
	}

	return flatStack, changed
}

func RemoveEmpty(tiles []Tile) []Tile {
	var flatStack []Tile
	for _, tile := range tiles {
		if !tile.empty {
			flatStack = append(flatStack, tile)
		}
	}
	return flatStack
}
