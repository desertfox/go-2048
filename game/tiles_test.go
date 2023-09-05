package game

import "testing"

func TestFlatten(t *testing.T) {
	cases := []struct {
		name     string
		start    []Tile
		expected []Tile
		changed  bool
	}{
		{
			name:     "Long shift",
			start:    []Tile{NewTile(2), NewTile(0), NewTile(0), NewTile(0)},
			expected: []Tile{NewTile(0), NewTile(0), NewTile(0), NewTile(2)},
			changed:  true,
		},
		{
			name:     "Long shift add",
			start:    []Tile{NewTile(2), NewTile(0), NewTile(0), NewTile(2)},
			expected: []Tile{NewTile(0), NewTile(0), NewTile(0), NewTile(4)},
			changed:  true,
		},
		{
			name:     "Long shift no add",
			start:    []Tile{NewTile(4), NewTile(0), NewTile(0), NewTile(2)},
			expected: []Tile{NewTile(0), NewTile(0), NewTile(4), NewTile(2)},
			changed:  true,
		},
		{
			name:     "Short shift",
			start:    []Tile{NewTile(0), NewTile(0), NewTile(2), NewTile(0)},
			expected: []Tile{NewTile(0), NewTile(0), NewTile(0), NewTile(2)},
			changed:  true,
		},
		{
			name:     "Short shift no add",
			start:    []Tile{NewTile(0), NewTile(4), NewTile(0), NewTile(2)},
			expected: []Tile{NewTile(0), NewTile(0), NewTile(4), NewTile(2)},
			changed:  true,
		},
		{
			name:     "Short shift add",
			start:    []Tile{NewTile(0), NewTile(0), NewTile(2), NewTile(2)},
			expected: []Tile{NewTile(0), NewTile(0), NewTile(0), NewTile(4)},
			changed:  true,
		},
		{
			name:     "add and shift",
			start:    []Tile{NewTile(2), NewTile(2), NewTile(0), NewTile(0)},
			expected: []Tile{NewTile(0), NewTile(0), NewTile(0), NewTile(4)},
			changed:  true,
		},
		{
			name:     "double shift no add",
			start:    []Tile{NewTile(2), NewTile(4), NewTile(0), NewTile(0)},
			expected: []Tile{NewTile(0), NewTile(0), NewTile(2), NewTile(4)},
			changed:  true,
		},
		{
			name:     "shift and add",
			start:    []Tile{NewTile(2), NewTile(2), NewTile(2), NewTile(2)},
			expected: []Tile{NewTile(0), NewTile(0), NewTile(4), NewTile(4)},
			changed:  true,
		},
		{
			name:     "shift add and not add",
			start:    []Tile{NewTile(2), NewTile(4), NewTile(2), NewTile(2)},
			expected: []Tile{NewTile(0), NewTile(2), NewTile(4), NewTile(4)},
			changed:  true,
		},
		{
			name:     "shift no change",
			start:    []Tile{NewTile(0), NewTile(0), NewTile(0), NewTile(2)},
			expected: []Tile{NewTile(0), NewTile(0), NewTile(0), NewTile(2)},
			changed:  false,
		},
		{
			name:     "full tile shift change on the right first",
			start:    []Tile{NewTile(2), NewTile(2), NewTile(2), NewTile(32)},
			expected: []Tile{NewTile(0), NewTile(2), NewTile(4), NewTile(32)},
			changed:  true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			shifted, changed := Flatten(tt.start, "right")

			if len(shifted) != len(tt.expected) {
				t.Errorf("wrong length: expected %v, but got %v", shifted, tt.expected)
				return
			}

			if changed != tt.changed {
				t.Errorf("Flatten didn't return changed got:%v, wanted:%v", changed, tt.changed)
			}

			for i, e := range tt.expected {
				if e.Value != shifted[i].Value {
					t.Errorf("expected %v, but got %v index:%d", e, shifted[i], i)
				}
			}

		})
	}
}
