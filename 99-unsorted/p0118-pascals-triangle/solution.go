// Brute Force
// Time complexity  : O(n^2)
// Space complexity : O(1)

package solution

func generate(numRows int) [][]int {
	rows := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if i == 0 {
				row[j] = 1
			} else if j == 0 {
				row[j] = rows[i-1][j]
			} else if j == i {
				row[j] = rows[i-1][j-1]
			} else {
				row[j] = rows[i-1][j-1] + rows[i-1][j]
			}
		}
		rows = append(rows, row)
	}

	return rows
}
