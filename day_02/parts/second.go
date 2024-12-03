package parts

import (
	"math"
	"os"
	"project/utils"
	"strconv"
	"strings"
)

func isReportSafe(nums []int) bool {
	if isSafeWithoutRemoval(nums) {
		return true
	}

	for i := 0; i < len(nums); i++ {
		reducedNums := make([]int, 0, len(nums)-1)
		reducedNums = append(reducedNums, nums[:i]...)
		reducedNums = append(reducedNums, nums[i+1:]...)

		if isSafeWithoutRemoval(reducedNums) {
			return true
		}
	}

	return false
}

func isSafeWithoutRemoval(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	increase := true
	decrease := true
	safe := true

	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]

		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			safe = false
			break
		}

		if nums[i] > nums[i+1] {
			increase = false
		}
		if nums[i] < nums[i+1] {
			decrease = false
		}
	}

	return safe && (increase || decrease)
}

func Second() {
	dat, err := os.ReadFile("data/input.txt")
	utils.Check(err)

	safeCounter := 0

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		split := strings.Split(line, " ")
		nums := make([]int, len(split))
		for i, s := range split {
			nums[i], err = strconv.Atoi(s)
			if err != nil {
				utils.Check(err)
			}
		}

		if isReportSafe(nums) {
			safeCounter++
		}
	}

	println(safeCounter)
}
