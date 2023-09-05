package game

import "testing"

func TestFlatten(t *testing.T) {
	cases := []struct {
		name     string
		start    []Tile
		expected []Tile
		changed  bool
	}{
		{"Long shift", []Tile{Tile(2), Tile(0), Tile(0), Tile(0)}, []Tile{Tile(0), Tile(0), Tile(0), Tile(2)}, true},
		{"Long shift add", []Tile{Tile(2), Tile(0), Tile(0), Tile(2)}, []Tile{Tile(0), Tile(0), Tile(0), Tile(4)}, true},
		{"Long shift no add", []Tile{Tile(4), Tile(0), Tile(0), Tile(2)}, []Tile{Tile(0), Tile(0), Tile(4), Tile(2)}, true},
		{"Short shift", []Tile{Tile(0), Tile(0), Tile(2), Tile(0)}, []Tile{Tile(0), Tile(0), Tile(0), Tile(2)}, true},
		{"Short shift no add", []Tile{Tile(0), Tile(4), Tile(0), Tile(2)}, []Tile{Tile(0), Tile(0), Tile(4), Tile(2)}, true},
		{"Short shift add", []Tile{Tile(0), Tile(0), Tile(2), Tile(2)}, []Tile{Tile(0), Tile(0), Tile(0), Tile(4)}, true},
		{"add and shift", []Tile{Tile(2), Tile(2), Tile(0), Tile(0)}, []Tile{Tile(0), Tile(0), Tile(0), Tile(4)}, true},
		{"double shift no add", []Tile{Tile(2), Tile(4), Tile(0), Tile(0)}, []Tile{Tile(0), Tile(0), Tile(2), Tile(4)}, true},
		{"shift and add", []Tile{Tile(2), Tile(2), Tile(2), Tile(2)}, []Tile{Tile(0), Tile(0), Tile(4), Tile(4)}, true},
		{"shift add and not add", []Tile{Tile(2), Tile(4), Tile(2), Tile(2)}, []Tile{Tile(0), Tile(2), Tile(4), Tile(4)}, true},
		{"shift no change", []Tile{Tile(0), Tile(0), Tile(0), Tile(2)}, []Tile{Tile(0), Tile(0), Tile(0), Tile(2)}, false},
		{"full tile shift change on the right first", []Tile{Tile(2), Tile(2), Tile(2), Tile(32)}, []Tile{Tile(0), Tile(2), Tile(4), Tile(32)}, true},
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
				if e != shifted[i] {
					t.Errorf("expected %v, but got %v index:%d", e, shifted[i], i)
				}
			}

		})
	}
}
