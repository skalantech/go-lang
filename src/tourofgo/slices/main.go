package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	result := make([][]uint8, dy)

	// for valy := 0; valy < dy; valy++
	for valy := range result {
		result[valy] = make([]uint8, dx)

		// for valx := 0; valx < dx; valx++
		for valx := range result[valy] {
			result[valy][valx] = uint8(valx^2*valy^3*valx*valy>>1) * 2
		}

	}
	return result
}

func main() {
	pic.Show(Pic)

	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
