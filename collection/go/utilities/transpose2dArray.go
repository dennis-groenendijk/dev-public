package main

// Transpose rearranges the position of given values from rows to columns
func Transpose(slice [][]string) [][]string {
	xAxis := len(slice[0])
	yAxis := len(slice)
	result := make([][]string, xAxis)

	for i := range result {
		result[i] = make([]string, yAxis)
	}
	for i := 0; i < xAxis; i++ {
		for j := 0; j < yAxis; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
