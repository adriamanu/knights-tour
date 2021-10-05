package main

import (
	"fmt"
	"strconv"
)

// generate bidimensional array that represent the chess board
func generateBoard(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "X"
		}
	}
	return board
}

// Knight can move using an L pattern
// Squares outside the chessboard are excluded
func getAvailableSquares(row int, col int, size int) [][]int {
	var availableSquares [][]int

	if row >= 2 && col < size-1 {
		// Up Right
		availableSquares = append(availableSquares, []int{row - 2, col + 1})
	}

	if row >= 2 && col >= 1 {
		// Up Left
		availableSquares = append(availableSquares, []int{row - 2, col - 1})
	}

	if row < size-1 && col < size-2 {
		// Right Down
		availableSquares = append(availableSquares, []int{row + 1, col + 2})
	}

	if row >= 1 && col < size-2 {
		// Right Up
		availableSquares = append(availableSquares, []int{row - 1, col + 2})
	}

	if row < size-1 && col >= 2 {
		// Left Down
		availableSquares = append(availableSquares, []int{row + 1, col - 2})
	}

	if row >= 1 && col >= 2 {
		// Left Up
		availableSquares = append(availableSquares, []int{row - 1, col - 2})
	}

	if row < size-2 && col < size-1 {
		// Down Right
		availableSquares = append(availableSquares, []int{row + 2, col + 1})
	}

	if row < size-2 && col >= 1 {
		// Down Left
		availableSquares = append(availableSquares, []int{row + 2, col - 1})
	}

	return availableSquares
}

func isSquareAlreadyVisited(board [][]string, row, col int) bool {
	return board[row][col] != "X"
}

func KnightTour(row int, col int, board [][]string, size int, count int) bool {
	board[row][col] = strconv.Itoa(count)

	if count == size*size {
		for r := range board {
			fmt.Println(board[r])
		}
		return true
	}

	// iterate over the possible moves
	availableSquares := getAvailableSquares(row, col, size)
	for _, square := range availableSquares {
		if !isSquareAlreadyVisited(board, square[0], square[1]) {
			if KnightTour(square[0], square[1], board, size, count+1) {
				return true
			}
		}
	}

	// backtrack
	board[row][col] = "X"
	return false
}

func main() {
	var size int = 6
	board := generateBoard(size)
	var count int = 1
	KnightTour(3, 2, board, size, count)
}
