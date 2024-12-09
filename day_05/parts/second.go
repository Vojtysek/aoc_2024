package parts

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Second() {
	// Read input file
	dat, err := os.ReadFile("data/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split input into ordering rules and update requests
	parts := strings.Split(string(dat), "\n\n")
	orderLines := strings.Split(parts[0], "\n")
	requestLines := strings.Split(parts[1], "\n")

	// Create map of ordering rules
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

		// Add or update ordering rule
		rule, exists := perfectOrder[i]
		if !exists {
			rule = OrderRule{page: i, before: []int{}}
		}
		rule.before = append(rule.before, b)
		perfectOrder[i] = rule
	}

	// Track sum of middle pages from correctly ordered updates
	var (
		correctOrderSum   int
		incorrectOrderSum int
	)

	// Check each update
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
			// Get middle page for correct order
			middleIndex := len(pageNumbers) / 2
			correctOrderSum += pageNumbers[middleIndex]
		} else {
			// Correct the order and get middle page
			correctedOrder := correctUpdateOrder(pageNumbers, perfectOrder)
			middleIndex := len(correctedOrder) / 2
			incorrectOrderSum += correctedOrder[middleIndex]
		}
	}

	fmt.Println("Sum of middle pages from correctly ordered updates:", correctOrderSum)
	fmt.Println("Sum of middle pages from corrected updates:", incorrectOrderSum)
}

// Correct the order of an update based on the rules
func correctUpdateOrder(pages []int, rules map[int]OrderRule) []int {
	// Create a copy of the pages slice to modify
	correctedPages := make([]int, len(pages))
	copy(correctedPages, pages)

	// Bubble sort with ordering rules
	for i := 0; i < len(correctedPages); i++ {
		for j := 0; j < len(correctedPages)-i-1; j++ {
			currentPage := correctedPages[j]
			nextPage := correctedPages[j+1]

			// Check if current page must come before next page
			rule, exists := rules[nextPage]
			if exists {
				// Check if current page should be after next page according to the rule
				if contains(rule.before, currentPage) {
					// Swap pages
					correctedPages[j], correctedPages[j+1] = correctedPages[j+1], correctedPages[j]
				}
			}
		}
	}

	return correctedPages
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	return false
}
