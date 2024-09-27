package day11

import "math"

const GALAXY_RUNE = '#'

type galaxyCluster struct {
	idToLocation                       map[int]location
	origRowGalaxyIDs, origColGalaxyIDs map[int][]int
	highestRow, highestCol, highestID  int
}

type location struct {
	row, col int
}

func (g *galaxyCluster) addGalaxy(loc location) {
	// Initialize maps if needed
	if g.origColGalaxyIDs == nil {
		g.origColGalaxyIDs = make(map[int][]int)
	}
	if g.origRowGalaxyIDs == nil {
		g.origRowGalaxyIDs = make(map[int][]int)
	}
	if g.idToLocation == nil {
		g.idToLocation = make(map[int]location)
	}

	// Add galaxy ID to row and column maps
	addToMap := func(m map[int][]int, key, id int) {
		m[key] = append(m[key], id)
	}
	addToMap(g.origColGalaxyIDs, loc.col, g.highestID)
	addToMap(g.origRowGalaxyIDs, loc.row, g.highestID)

	// Update location and highest ID
	g.idToLocation[g.highestID] = loc
	g.highestID++

	// Update highest column and row
	if loc.col > g.highestCol {
		g.highestCol = loc.col
	}
	if loc.row > g.highestRow {
		g.highestRow = loc.row
	}
}

// Refactor by ChatGPT
func (g *galaxyCluster) expandBy(expandAmount int) {
	for _, data := range []struct {
		IDs               *map[int][]int
		limit             int
		getRowOrColumnPtr func(loc *location) *int
	}{
		{
			IDs:   &g.origRowGalaxyIDs,
			limit: g.highestRow,
			getRowOrColumnPtr: func(loc *location) *int {
				return &loc.row
			},
		},
		{
			IDs:   &g.origColGalaxyIDs,
			limit: g.highestCol,
			getRowOrColumnPtr: func(loc *location) *int {
				return &loc.col
			},
		},
	} {
		for i := 0; i < data.limit; i++ {
			_, hasGalaxy := (*data.IDs)[i]
			if !hasGalaxy {
				for idx, IDs := range *data.IDs {
					if idx > i {
						for _, id := range IDs {
							loc := g.idToLocation[id]
							*data.getRowOrColumnPtr(&loc) += expandAmount
							g.idToLocation[id] = loc
						}
					}
				}
			}
		}
	}
}

func Absi(number int) int {
	return int(math.Abs(float64(number)))
}

func (gc *galaxyCluster) getDistanceBetweenGalaxies(ida, idb int) int {
	coldiff := gc.idToLocation[ida].col - gc.idToLocation[idb].col
	rowdiff := gc.idToLocation[ida].row - gc.idToLocation[idb].row
	return Absi(coldiff) + Absi(rowdiff)
}

func (gc *galaxyCluster) makeAllPossiblePairs() [][]int {
	pairs := [][]int{}
	for ida := 0; ida < len(gc.idToLocation)-1; ida++ {
		for idb := ida + 1; idb < len(gc.idToLocation); idb++ {
			pairs = append(pairs, []int{ida, idb})
		}
	}
	return pairs
}

func (gc *galaxyCluster) getAllDistances() []int {
	var distances []int
	for _, pair := range gc.makeAllPossiblePairs() {
		distances = append(distances, gc.getDistanceBetweenGalaxies(pair[0], pair[1]))
	}
	return distances
}
