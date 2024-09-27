package day11

import "math"

const GALAXY_RUNE = '#'

type galaxyCluster struct {
	idToLocation                      map[int]location
	rowGalaxyIDs, colGalaxyIDs        map[int][]int
	highestRow, highestCol, highestID int
}

type location struct {
	row, col int
}

func (g *galaxyCluster) addGalaxy(loc location) {
	if g.colGalaxyIDs == nil {
		g.colGalaxyIDs = make(map[int][]int)
	}
	g.colGalaxyIDs[loc.col] = append(g.colGalaxyIDs[loc.col], g.highestID)

	if g.rowGalaxyIDs == nil {
		g.rowGalaxyIDs = make(map[int][]int)
	}
	g.rowGalaxyIDs[loc.row] = append(g.rowGalaxyIDs[loc.row], g.highestID)

	if g.idToLocation == nil {
		g.idToLocation = make(map[int]location)
	}
	g.idToLocation[g.highestID] = loc
	g.highestID++

	if loc.col > g.highestCol {
		g.highestCol = loc.col
	}

	if loc.row > g.highestRow {
		g.highestRow = loc.row
	}
}

func (g *galaxyCluster) expand() {
	for i := 0; i < g.highestRow; i++ {
		_, ok := g.rowGalaxyIDs[i]
		if !ok {
			// row i is empty
			for rowIdx, IDs := range g.rowGalaxyIDs {
				if rowIdx > i {
					for _, id := range IDs {
						g.idToLocation[id] = location{
							row: g.idToLocation[id].row + 1,
							col: g.idToLocation[id].col,
						}
					}
				}
			}
		}
	}

	for i := 0; i < g.highestCol; i++ {
		_, ok := g.colGalaxyIDs[i]
		if !ok {
			// col i is empty
			for colIdx, IDs := range g.colGalaxyIDs {
				if colIdx > i {
					for _, id := range IDs {
						g.idToLocation[id] = location{
							col: g.idToLocation[id].col + 1,
							row: g.idToLocation[id].row,
						}
					}
				}
			}
		}
	}
}

func intabs(number int) int {
	return int(math.Abs(float64(number)))
}

func (gc *galaxyCluster) getDistanceBetweenGalaxies(ida, idb int) int {
	coldiff := gc.idToLocation[ida].col - gc.idToLocation[idb].col
	rowdiff := gc.idToLocation[ida].row - gc.idToLocation[idb].row
	return intabs(coldiff) + intabs(rowdiff)
}
