package days

import (
	"fmt"
	"slices"
	"strings"
)

func cloneGrid(src [][]byte) [][]byte {
	dst := make([][]byte, len(src))
	for i := range src {
		dst[i] = make([]byte, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func checkNeighbours(inputGrid [][]byte) (int, [][]byte) {
	ogGrid := cloneGrid(inputGrid)
	nRows := len(inputGrid) - 1
	nCols := len(inputGrid[0]) - 1

	// var top, bottom, left, right, topLeft, topRight, bottomLeft, bottomRight bool

	lessThan4Neightbours := 0
	for row, levels := range inputGrid {
		for col, ch := range levels {
			totNeighbours := 0

			if ch == '@' {
				// Top
				if row != 0 && rune(inputGrid[row-1][col]) == '@' {
					totNeighbours++
				}

				// Bottom
				if row != nRows && rune(inputGrid[row+1][col]) == '@' {
					totNeighbours++
				}

				// Left
				if col != 0 && rune(inputGrid[row][col-1]) == '@' {
					totNeighbours++
				}

				// Right
				if col != nCols && rune(inputGrid[row][col+1]) == '@' {
					totNeighbours++
				}

				// TopLeft
				if row != 0 && col != 0 && rune(inputGrid[row-1][col-1]) == '@' {
					totNeighbours++
				}

				// TopRight
				if row != 0 && col != nCols && rune(inputGrid[row-1][col+1]) == '@' {
					totNeighbours++
				}

				// BottomLeft
				if row != nRows && col != 0 && rune(inputGrid[row+1][col-1]) == '@' {
					totNeighbours++
				}

				// BottomRight
				if row != nRows && col != nCols && rune(inputGrid[row+1][col+1]) == '@' {
					totNeighbours++
				}
				if totNeighbours < 4 {
					lessThan4Neightbours++
					ogGrid[row][col] = 'x'
				}
			}

			// Checking for less than 4 neighbours

			// Resetting every valid @ with x

		}
	}

	return lessThan4Neightbours, ogGrid
}

func D4P1(inputFile string) {
	inputGrid := slices.Collect(strings.SplitSeq(inputFile, "\n"))
	validRollPapesAsByes := make([][]byte, len(inputGrid))

	for i := range inputGrid {
		validRollPapesAsByes[i] = []byte(inputGrid[i])
	}

	validRollPapes, validRollPapesAsByes := checkNeighbours(validRollPapesAsByes)

	fmt.Println("Day 4, Part 1: ", validRollPapes)
}

func D4P2(inputFile string) {
	inputGrid := slices.Collect(strings.SplitSeq(inputFile, "\n"))
	validRollPapesAsByes := make([][]byte, len(inputGrid))
	iterativeCheck := 0

	validRollPapes := -1
	for i := range inputGrid {
		validRollPapesAsByes[i] = []byte(inputGrid[i])
	}

	for validRollPapes != 0 {
		validRollPapes, validRollPapesAsByes = checkNeighbours(validRollPapesAsByes)
		iterativeCheck += validRollPapes
	}

	fmt.Println("Day 4, Part 2: ", iterativeCheck)
}
