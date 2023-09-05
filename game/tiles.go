package game

type Tile int

func (t *Tile) Double() {
	if *t == 0 {
		*t = 2
		return
	}
	*t = *t * 2
}

func Flatten(tiles []Tile, direction string) ([]Tile, bool) {
	flatStack := RemoveEmpty(tiles)

	switch direction {
	case "right":
		for i := len(flatStack) - 1; i > 0; i-- {
			if flatStack[i] == flatStack[i-1] {
				flatStack[i].Double()
				flatStack[i-1] = Tile(0)
			}
		}
	case "left":
		for i := 0; len(flatStack)-1 > i; i++ {
			if flatStack[i] == flatStack[i+1] {
				flatStack[i].Double()
				flatStack[i+1] = Tile(0)
			}
		}
	}

	flatStack = RemoveEmpty(flatStack)

	for len(flatStack) < 4 {
		if direction == "right" {
			flatStack = append([]Tile{Tile(0)}, flatStack...)
		} else {
			flatStack = append(flatStack, Tile(0))
		}
	}

	for i := 0; i < 4; i++ {
		if tiles[i] != flatStack[i] {
			return flatStack, true
		}
	}

	return flatStack, false
}

func RemoveEmpty(tiles []Tile) []Tile {
	var flatStack []Tile
	for _, tile := range tiles {
		if tile != 0 {
			flatStack = append(flatStack, tile)
		}
	}
	return flatStack
}
