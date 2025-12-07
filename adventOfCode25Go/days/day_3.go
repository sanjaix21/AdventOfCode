package days

import (
	"fmt"
	"strconv"
	"strings"
)

// Function to reset array elements to 0 starting from x index
// Example:
// output = reset([1, 2, 3, 4], 2)
// output: [1, 2, 0, 0]
func reset(arr []int, from int) {
	for i := from; i < len(arr); i++ {
		arr[i] = 0
	}
}

// Function to convert int array to a single number
func formNumber(arr []int) int {
	result := 0
	for _, num := range arr {
		result = result*10 + num
	}
	return result
}

// Function to find largest number possible in an array in left to right direction
func findLargestJolt(jolts []int, number, rem int, size int) {
	totalElement := size
	elementLeft := size - 1
	for idx := range totalElement {
		if number > jolts[idx] && rem >= elementLeft {
			jolts[idx] = number
			reset(jolts, idx+1)
			break
		}
		elementLeft--
	}
}

func D3P1(inputFile string) {
	totOutputJoltage := 0
	for bank := range strings.SplitSeq(inputFile, "\n") {
		jolts := make([]int, 2)
		for i := range len(bank) {
			number, _ := strconv.Atoi(string(bank[i]))
			remaining := len(bank) - i - 1
			findLargestJolt(jolts, number, remaining, 2)
		}

		largestJoltInBank := formNumber(jolts)
		totOutputJoltage += largestJoltInBank
	}

	fmt.Println("Day 3, Part 1: ", totOutputJoltage)
}

func D3P2(inputFile string) {
	totOutputJoltage12Dig := 0

	for bank := range strings.SplitSeq(inputFile, "\n") {
		jolts := make([]int, 12)

		for i := range len(bank) {
			number, _ := strconv.Atoi(string(bank[i]))
			remaining := len(bank) - i - 1

			findLargestJolt(jolts, number, remaining, 12)
		}

		largestJoltInBank := formNumber(jolts)
		totOutputJoltage12Dig += largestJoltInBank
	}

	fmt.Println("Day 3, Part 2: ", totOutputJoltage12Dig)
}
