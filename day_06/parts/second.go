package parts

import (
	"fmt"
	"os"
	"strings"
)

func simulateGuardPath(mapOfGame [][]rune, initialState State) bool {
	currentState := initialState
	visitedStates := make(map[string]int)

	directions := [][]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	for step := 0; step < len(mapOfGame)*len(mapOfGame[0])*4; step++ {
		// Create a unique key for the current state
		stateKey := fmt.Sprintf("%d,%d,%d", currentState.row, currentState.col, currentState.direction)

		// If we've seen this exact state before, we're in a loop
		if visitedStates[stateKey] > 0 {
			return true
		}
		visitedStates[stateKey]++

		nextRow := currentState.row + directions[currentState.direction][0]
		nextColumn := currentState.col + directions[currentState.direction][1]

		// Check if out of bounds
		if nextRow < 0 || nextRow >= len(mapOfGame) ||
			nextColumn < 0 || nextColumn >= len(mapOfGame[0]) {
			return false
		}

		// Hit a wall
		if mapOfGame[nextRow][nextColumn] == '#' {
			currentState.direction = (currentState.direction + 1) % 4
		} else {
			// Move to the next position
			currentState.row = nextRow
			currentState.col = nextColumn
		}
	}

	return false
}

func countLoopPositions(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mapOfGame := make([][]rune, len(lines))
	var initialState State

	// Parse initial map and find starting position
	for rowIndex, line := range lines {
		mapOfGame[rowIndex] = []rune(line)
		for colIndex, char := range mapOfGame[rowIndex] {
			if char == '^' {
				initialState = State{
					row:       rowIndex,
					col:       colIndex,
					direction: 0, // Up
				}
			}
		}
	}

	loopPositions := 0

	// Try placing obstruction at every possible empty position
	for rowIndex := 0; rowIndex < len(mapOfGame); rowIndex++ {
		for colIndex := 0; colIndex < len(mapOfGame[0]); colIndex++ {
			// Skip the starting position
			if rowIndex == initialState.row && colIndex == initialState.col {
				continue
			}

			// Skip existing walls
			if mapOfGame[rowIndex][colIndex] == '#' {
				continue
			}

			// Create a copy of the map and place an obstruction
			modifiedMap := make([][]rune, len(mapOfGame))
			for i := range mapOfGame {
				modifiedMap[i] = make([]rune, len(mapOfGame[i]))
				copy(modifiedMap[i], mapOfGame[i])
			}
			modifiedMap[rowIndex][colIndex] = '#'

			// Simulate guard's path with the new obstruction
			if simulateGuardPath(modifiedMap, initialState) {
				loopPositions++
			}
		}
	}

	return loopPositions
}

func Second() {
	dat, err := os.ReadFile("data/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := countLoopPositions(string(dat))
	fmt.Println("Number of positions that cause the guard to get stuck:", result)
}
