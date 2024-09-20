package day10

import (
	"fmt"
	"os"
	"strings"
)

const (
	START_TILE = 'S'
)

type pipeSystem struct {
	system []string
}

func (s *pipeSystem) toHtml() string {
	html := `<html><head><link rel="stylesheet" href="day10.css"></head><body style="background-color: #333;"><table><tbody>`
	for rowIdx, line := range s.system {
		html += `<tr>`
		for colIdx, r := range line {
			html += fmt.Sprintf("<td id=\"loc-%d-%d\">%s</td>", rowIdx, colIdx, string(r))
		}
		html += `</tr>`
	}
	html += `</tbody></table></body></html>`
	return html
}

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

func (s *pipeSystem) getNeighbor(row, col int, directions string) string {
	var neighbors string = ""
	for _, d := range directions {
		nRow, nCol := directionToCoordinate(row, col, string(d))
		neighbors += string(s.system[nRow][nCol])
	}
	return neighbors
}

var (
	connections = map[string]string{
		"-": "EW",
		"|": "NS",
		"7": "WS",
		"J": "NW",
		"L": "EN",
		"F": "SE",
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

func (s *pipeSystem) calcLoopLength(row, col int) int {
	var oldRow, oldCol int
	var css string
	css += `td {
	color: gray
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
		for _, r := range connections[currentTile] {
			if s.canGo(row, col, string(r)) {
				// fmt.Printf("At %d, %d, '%s' - going %s\n", row, col, string(currentTile), string(r))
				oldCol = col
				oldRow = row
				row, col = directionToCoordinate(row, col, string(r))
				rowText := []rune(s.system[oldRow])
				rowText[oldCol] = '.'
				s.system[oldRow] = string(rowText)
				moved = true
				css += fmt.Sprintf(`#loc-%d-%d {
    color: #%03x;
	animation-name: example;
	animation-duration: %dms;
}
`, row, col, i+16, i*10)
				break
			} else {
				// fmt.Printf("At %d, %d, '%s' - can't go %s\n", row, col, string(currentTile), string(r))
			}
		}
		if !moved {
			os.WriteFile("day10/day10.css", []byte(css), 0644)
			return i
		}
	}
	return -1
}
