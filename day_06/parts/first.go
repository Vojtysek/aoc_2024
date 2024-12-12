package parts

import (
	"fmt"
	"os"
	"strings"
)

type State struct {
	row       int
	col       int
	direction int
}

func solveGuardPath(input string) (int, [][]rune) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mapOfGame := make([][]rune, len(lines))
	var initialState State

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

	visitedPositions := make(map[string]bool)
	currentState := initialState

	directions := [][]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	visitedPositions[fmt.Sprintf("%d,%d", currentState.row, currentState.col)] = true

	for {
		nextRow := currentState.row + directions[currentState.direction][0]
		nextColumn := currentState.col + directions[currentState.direction][1]

		if nextRow < 0 || nextRow >= len(mapOfGame) ||
			nextColumn < 0 || nextColumn >= len(mapOfGame[0]) {
			break
		}

		if mapOfGame[nextRow][nextColumn] == '#' {
			currentState.direction = (currentState.direction + 1) % 4
		} else {
			currentState.row = nextRow
			currentState.col = nextColumn

			posKey := fmt.Sprintf("%d,%d", currentState.row, currentState.col)
			if !visitedPositions[posKey] {
				visitedPositions[posKey] = true
			}
		}
	}

	return len(visitedPositions), mapOfGame
}

func First() {
	dat, err := os.ReadFile("data/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	distinctPositions, _ := solveGuardPath(string(dat))
	fmt.Println("Distinct positions:", distinctPositions)
}
