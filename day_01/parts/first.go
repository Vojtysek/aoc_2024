package parts

import (
	"math"
	"os"
	"project/utils"
	"sort"
	"strconv"
	"strings"
)

func First() {

	dat, err := os.ReadFile("data/input.txt")
	utils.Check(err)

	lines := strings.Split(string(dat), "\n")

	var left []int
	var right []int

	for _, line := range lines {
		split := strings.Split(line, "   ")
		leftNumber := split[0]
		rightNumber := split[1]

		l, err := strconv.Atoi(leftNumber)
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(rightNumber)
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	var distance int = 0

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i]) - float64(right[i])))
	}

	println(int(distance))

}
