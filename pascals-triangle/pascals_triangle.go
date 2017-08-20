package pascal

const testVersion = 1

// Triangle returns nr rows of Pascal's triangle
func Triangle(nr int) [][]int {
	rows := make([][]int, nr)

	rows[0] = []int{1}

	for row := 1; row < nr; row++ {

		r := make([]int, len(rows[row-1]))
		copy(r, rows[row-1])

		for i := 1; i < len(rows[row-1]); i++ {
			r = append(r[:i], rows[row-1][i-1]+rows[row-1][i])
		}
		r = append(r, 1)

		rows[row] = r

	}
	return rows
}
