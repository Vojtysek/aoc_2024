package parts

import (
	"os"
	"project/utils"
	"sort"
	"strconv"
	"strings"
)

func Second() {

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

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var distance int = 0

	for i := 0; i < len(left); i++ {
		/* find how many times is left number in right */
		count := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				count++
			}
		}

		distance += left[i] * count

	}

	println(distance)

}
