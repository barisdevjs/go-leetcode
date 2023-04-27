package main

func Generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := -1; j < i; j++ {
			row[j+1] = (j + 2) * i
		}
		result[i] = row
	}
	return result
}
