package day09

// Takes a sequence of ints and returns the next int in the sequence
func extrapolate(sequence []int) int {
	diffs := [][]int{
		sequence,
	}
	diffs = append(diffs, calcDifferencesRecurse(sequence)...)
	for i := len(diffs) - 1; i > 0; i-- {
		currentRow := diffs[i]
		nextRow := diffs[i-1]
		currentRowLastNumber := currentRow[len(currentRow)-1]
		nextRowLastNumber := nextRow[len(nextRow)-1]
		diffs[i-1] = append(nextRow, currentRowLastNumber+nextRowLastNumber)
	}
	return diffs[0][len(diffs[0])-1]
}

// Takes a sequence of ints and returns the int before the first in the sequence
func extrapolatePrev(sequence []int) int {
	diffs := [][]int{
		sequence,
	}
	diffs = append(diffs, calcDifferencesRecurse(sequence)...)
	for i := len(diffs) - 1; i > 0; i-- {
		currentRow := diffs[i]
		nextRow := diffs[i-1]
		currentRowFirstNumber := currentRow[0]
		nextRowFirstNumber := nextRow[0]
		diffs[i-1] = append([]int{nextRowFirstNumber - currentRowFirstNumber}, nextRow...)
	}
	return diffs[0][0]
}

// Returns a new array, consisting of the amount by which the numbers in the
// input array increase. E.g. [2,4,6] -> [2,2]
func calcDifferences(sequence []int) []int {
	var diffs []int
	for i, s := range sequence {
		if i > 0 {
			diffs = append(diffs, s-sequence[i-1])
		}
	}
	return diffs
}

func calcDifferencesRecurse(sequence []int) [][]int {
	result := [][]int{}
	diffs := calcDifferences(sequence)
	result = append(result, diffs)
	for _, d := range diffs {
		if d != 0 {
			result = append(result, calcDifferencesRecurse(diffs)...)
			break
		}
	}
	return result
}
