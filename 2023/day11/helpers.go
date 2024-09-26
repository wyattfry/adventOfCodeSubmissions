package day11

type galaxy struct {
	row, col int
}

const GALAXY_RUNE = '#'

func expand(gals []galaxy) {
	var maxRow, maxCol int
	for _, g := range gals {
		if g.col > maxCol {
			maxCol = g.col
		}
		if g.row > maxRow {
			maxRow = g.row
		}
	}
	for r := range maxRow {
		for c := range maxCol {
			
		}
	}
}
