package parts

import (
	"os"
	"project/utils"
	"strings"
)

var xmasCounterPuzzle int

func Second() {

	dat, err := os.ReadFile("data/input.txt")
	utils.Check(err)
	lines := strings.Split(string(dat), "\n")

	var charMatrix [][]rune

	for _, line := range lines {
		charMatrix = append(charMatrix, []rune(line))
	}

	for i := 1; i < len(charMatrix)-1; i++ {
		for j := 1; j < len(charMatrix[i])-1; j++ {
			if charMatrix[i][j] == 'A' {
				LT := charMatrix[i-1][j-1]
				LB := charMatrix[i+1][j-1]
				RT := charMatrix[i-1][j+1]
				RB := charMatrix[i+1][j+1]

				countOfM := 0
				countOfS := 0

				letters := []rune{LT, LB, RT, RB}

				for _, letter := range letters {
					if letter == 'M' {
						countOfM++
					} else if letter == 'S' {
						countOfS++
					}
				}

				if countOfM == 2 && countOfS == 2 {
					if (LT == LB && RT == RB) || (LT == RT && LB == RB) {
						xmasCounterPuzzle++
					}
				}
			}
		}
	}

	println(xmasCounterPuzzle)
}
