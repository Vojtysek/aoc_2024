package parts

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OrderRule struct {
	page   int
	before []int
}

func First() {
	dat, err := os.ReadFile("data/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	parts := strings.Split(string(dat), "\n\n")
	orderLines := strings.Split(parts[0], "\n")
	requestLines := strings.Split(parts[1], "\n")

	perfectOrder := make(map[int]OrderRule)
	for _, order := range orderLines {
		if order == "" {
			continue
		}
		splitOrder := strings.Split(order, "|")
		b, err := strconv.Atoi(splitOrder[0])
		if err != nil {
			fmt.Println("Error converting before page:", err)
			continue
		}
		i, err := strconv.Atoi(splitOrder[1])
		if err != nil {
			fmt.Println("Error converting index page:", err)
			continue
		}

		rule, exists := perfectOrder[i]
		if !exists {
			rule = OrderRule{page: i, before: []int{}}
		}
		rule.before = append(rule.before, b)
		perfectOrder[i] = rule
	}

	var middlePageSum int

	for _, request := range requestLines {
		if request == "" {
			continue
		}
		splitRequest := strings.Split(request, ",")
		pageNumbers := make([]int, len(splitRequest))
		for i, pageStr := range splitRequest {
			page, err := strconv.Atoi(pageStr)
			if err != nil {
				fmt.Println("Error converting page:", err)
				break
			}
			pageNumbers[i] = page
		}

		// Check if the update is in correct order
		if isCorrectOrder(pageNumbers, perfectOrder) {
			// Get middle page
			middleIndex := len(pageNumbers) / 2
			middlePageSum += pageNumbers[middleIndex]
		}
	}

	fmt.Println(middlePageSum)
}

func isCorrectOrder(pages []int, rules map[int]OrderRule) bool {
	for i := 0; i < len(pages); i++ {
		currentPage := pages[i]

		rule, exists := rules[currentPage]
		if !exists {
			continue
		}

		for _, beforePage := range rule.before {
			beforeIndex := findPageIndex(pages, beforePage)
			currentIndex := findPageIndex(pages, currentPage)

			if beforeIndex != -1 && currentIndex != -1 && beforeIndex > currentIndex {
				return false
			}
		}
	}
	return true
}

func findPageIndex(pages []int, page int) int {
	for i, p := range pages {
		if p == page {
			return i
		}
	}
	return -1
}
