package parts

import (
	"os"
	"project/utils"
	"strings"
)

var xmasCounter int

func findHorizontal(charMatrix [][]rune, i, j int) {
	if j+3 < len(charMatrix[i]) &&
		charMatrix[i][j+1] == 'M' &&
		charMatrix[i][j+2] == 'A' &&
		charMatrix[i][j+3] == 'S' {
		xmasCounter++
	}

	if j-3 >= 0 &&
		charMatrix[i][j-1] == 'M' &&
		charMatrix[i][j-2] == 'A' &&
		charMatrix[i][j-3] == 'S' {
		xmasCounter++
	}
}

func findVertical(charMatrix [][]rune, i, j int) {
	if i+3 < len(charMatrix) &&
		charMatrix[i+1][j] == 'M' &&
		charMatrix[i+2][j] == 'A' &&
		charMatrix[i+3][j] == 'S' {
		xmasCounter++
	}

	if i-3 >= 0 &&
		charMatrix[i-1][j] == 'M' &&
		charMatrix[i-2][j] == 'A' &&
		charMatrix[i-3][j] == 'S' {
		xmasCounter++
	}
}

func findDiagonal(charMatrix [][]rune, i, j int) {
	if i+3 < len(charMatrix) && j+3 < len(charMatrix[i]) &&
		charMatrix[i+1][j+1] == 'M' &&
		charMatrix[i+2][j+2] == 'A' &&
		charMatrix[i+3][j+3] == 'S' {
		xmasCounter++
	}

	if i-3 >= 0 && j-3 >= 0 &&
		charMatrix[i-1][j-1] == 'M' &&
		charMatrix[i-2][j-2] == 'A' &&
		charMatrix[i-3][j-3] == 'S' {
		xmasCounter++
	}

	if i+3 < len(charMatrix) && j-3 >= 0 &&
		charMatrix[i+1][j-1] == 'M' &&
		charMatrix[i+2][j-2] == 'A' &&
		charMatrix[i+3][j-3] == 'S' {
		xmasCounter++
	}

	if i-3 >= 0 && j+3 < len(charMatrix[i]) &&
		charMatrix[i-1][j+1] == 'M' &&
		charMatrix[i-2][j+2] == 'A' &&
		charMatrix[i-3][j+3] == 'S' {
		xmasCounter++
	}
}

func First() {
	dat, err := os.ReadFile("data/input.txt")
	utils.Check(err)
	lines := strings.Split(string(dat), "\n")

	var charMatrix [][]rune

	for _, line := range lines {
		charMatrix = append(charMatrix, []rune(line))
	}

	for i := 0; i < len(charMatrix); i++ {
		for j := 0; j < len(charMatrix[i]); j++ {
			if charMatrix[i][j] == 'X' {
				findHorizontal(charMatrix, i, j)
				findVertical(charMatrix, i, j)
				findDiagonal(charMatrix, i, j)
			}
		}
	}
	println(xmasCounter)

}
