package day10

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	START_TILE = 'S'
)

type pipeSystem struct {
	system        []string
	startTileType string
}

func (s *pipeSystem) toHtml() string {
	toPipe := map[rune]rune{
		'-': '━',
		'|': '┃',
		'L': '┗',
		'F': '┏',
		'7': '┓',
		'J': '┛',
		'S': '╋',
	}
	html := `<html><head><link rel="stylesheet" href="day10.css"></head><body style="background-color: #333;"><table><tbody>`
	for rowIdx, line := range s.system {
		html += `<tr>`
		for colIdx, r := range line {
			html += fmt.Sprintf("<td id=\"loc-%d-%d\">%s</td>", rowIdx, colIdx, string(toPipe[r]))
		}
		html += `</tr>`
	}
	html += `</tbody></table></body></html>`
	return html
}

// Find the index of the first tile of the given type, useful for finding the starting tile.
func (s *pipeSystem) indexOf(r rune) (row, col int) {
	for rowIdx, row := range s.system {
		for colIdx, col := range row {
			if col == r {
				return rowIdx, colIdx
			}
		}
	}
	return -1, -1
}

// Get coordinate of tile in a given direction from a starting tile
func directionToCoordinate(row, col int, direction string) (int, int) {
	switch direction {
	case "N":
		return row - 1, col + 0
	case "E":
		return row - 0, col + 1
	case "S":
		return row + 1, col + 0
	case "W":
		return row - 0, col - 1
	}
	return -1, -1
}

// Get types of adjacent tiles in the given direction(s)
func (s *pipeSystem) getNeighbor(row, col int, directions string) string {
	var neighbors string = ""
	for _, d := range directions {
		nRow, nCol := directionToCoordinate(row, col, string(d))
		if nCol >= 0 && nCol < len(s.system[0]) && nRow >= 0 && nRow < len(s.system) {
			neighbors += string(s.system[nRow][nCol])
		}
	}
	return neighbors
}

var (
	connections = map[string]string{
		"-": "EW",
		"|": "NS",
		"7": "SW",
		"J": "NW",
		"L": "NE",
		"F": "ES",
		"S": "NSEW",
	}

	opposites = map[string]string{
		"W": "E",
		"E": "W",
		"N": "S",
		"S": "N",
	}
)

func (s *pipeSystem) canGo(row, col int, direction string) bool {
	opposite := opposites[direction]
	return strings.Contains(connections[s.getNeighbor(row, col, direction)], opposite)
}

// This function is too long.
func (s *pipeSystem) calcLoopLength(row, col int) int {
	var oldRow, oldCol int
	var css string
	css += `td {
	color: gray;
	padding: 0;
    font-family: monospace;
    line-height: 0.4rem;
}
@keyframes example {
  from {color: red;}
  to {color: yellow;}
}
`
	for i := 0; true; i++ {
		currentTile := string(s.system[row][col])
		var moved bool
		if currentTile == string(START_TILE) {
			// figure out actual tile type
			var canGoDirs string
			for _, d := range "NESW" {
				if s.canGo(row, col, string(d)) {
					canGoDirs += string(d)
				}
			}
			for k, v := range connections {
				if canGoDirs == v {
					s.startTileType = k
					break
				}
			}
		}
		for _, r := range connections[currentTile] {
			if s.canGo(row, col, string(r)) {
				oldCol = col
				oldRow = row
				row, col = directionToCoordinate(row, col, string(r))
				rowText := []rune(s.system[oldRow])
				rowText[oldCol] = '.'
				s.system[oldRow] = string(rowText)
				moved = true
				color := int(math.Abs(math.Sin(float64(i)/1000) * 0xfff))
				css += fmt.Sprintf(`#loc-%d-%d {
    color: #%03x;
	animation-name: example;
	animation-duration: %dms;
}
`, row, col, color, i*10)
				break
			}
		}
		if !moved {
			os.WriteFile("day10/day10.css", []byte(css), 0644)
			return i
		}
	}
	return -1
}

////////// Part 2 //////////////

type location struct {
	row, col int
}

type tile struct {
	loc          location
	isPartOfLoop bool
	tileType     string
	north        *tile
	east         *tile
	south        *tile
	west         *tile
}

func (t *tile) getNeighbor(direction rune) *tile {
	switch direction {
	case 'N':
		return t.north
	case 'E':
		return t.east
	case 'S':
		return t.south
	case 'W':
		return t.west
	}
	return nil
}

// Infer tile type based on neighboring tiles. Useful for the start tile.
func inferType(loc location, lines []string) string {
	var canGoDirs string
	ps := pipeSystem{
		system: lines,
	}
	for _, d := range "NESW" {
		if ps.canGo(loc.row, loc.col, string(d)) {
			canGoDirs += string(d)
		}
	}
	for k, v := range connections {
		if canGoDirs == v {
			return k
		}
	}
	return ""
}

func parseInput(lines []string) map[string]*tile {
	var tiles = map[string]*tile{}
	for rowIdx, row := range lines {
		for colIdx, col := range row {
			currentTile := tile{
				loc: location{
					row: rowIdx,
					col: colIdx,
				},
				tileType: string(col),
			}
			tiles[fmt.Sprintf("%d,%d", rowIdx, colIdx)] = &currentTile

			// Make start tile easy to find
			if col == START_TILE {
				currentTile.isPartOfLoop = true
				tiles[string(START_TILE)] = &currentTile
				currentTile.tileType = inferType(currentTile.loc, lines)
			}

			north, ok := tiles[fmt.Sprintf("%d,%d", rowIdx-1, colIdx+0)]
			if ok {
				currentTile.north = north
				north.south = &currentTile
			}

			east, ok := tiles[fmt.Sprintf("%d,%d", rowIdx-0, colIdx+1)]
			if ok {
				currentTile.east = east
				east.west = &currentTile
			}

			south, ok := tiles[fmt.Sprintf("%d,%d", rowIdx+1, colIdx+0)]
			if ok {
				currentTile.south = south
				south.north = &currentTile
			}

			west, ok := tiles[fmt.Sprintf("%d,%d", rowIdx-0, colIdx-1)]
			if ok {
				currentTile.west = west
				west.east = &currentTile
			}
		}
	}
	return tiles
}

// Traverse the loop, setting each tile's `isPartOfLoop` to `true`
func (t *tile) traceLoop() []*tile {
	// Infinite loop until path returns to start
	var pathAsSlice []*tile
	for i := 0; true; i++ {
		var moved bool
		pathAsSlice = append(pathAsSlice, t)
		for _, r := range connections[t.tileType] {
			neighbor := t.getNeighbor(r)
			if neighbor != nil && !neighbor.isPartOfLoop {
				neighbor.isPartOfLoop = true
				t = neighbor
				moved = true
				break
			}
		}
		if !moved {
			break
		}
	}
	return pathAsSlice
}

// Find the area of the polygon via shoelace method
func calculateAreaShoelace(pathSlice []*tile) int {
	var area int
	for i, t := range pathSlice {
		nxt := (i + 1) % len(pathSlice)
		x_i := t.loc.col
		y_i := t.loc.row
		y_ip1 := pathSlice[nxt].loc.row
		x_ip1 := pathSlice[nxt].loc.col
		thisStep := x_i*y_ip1 - x_ip1*y_i
		area += thisStep
	}
	area /= 2
	return area
}

// Find the number of interior points using Pick's Theorem
func calculateInteriorPointsPicks(area int, boundaryPointCount int) int {
	return area + 1 - boundaryPointCount/2
}
