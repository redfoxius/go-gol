package service

import "fmt"

func DrawMatrix(input [][]int) {
	for _, row := range input {
		fmt.Println(generateRow(row))
	}
	fmt.Println()
	fmt.Println(`Press Ctrl+C to terminate execution`)
}

func generateRow(row []int) string {
	res := `|`

	for _, v := range row {
		if v == 1 {
			res += `#|`
		} else {
			res += ` |`
		}
	}

	return res
}
