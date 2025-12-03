package main

import (
	"aoc25/utils"
	"fmt"
	"strconv"
	"strings"
)

func reset(arr []int, from int) {
	for i := from; i < len(arr); i++ {
		arr[i] = 0
	}
	fmt.Println("Array after reset: ", arr)
}

func formNumber(arr []int) int {
	result := 0
	fmt.Println("jolts came as: ", arr)
	for _, num := range arr {
		result = result*10 + num
	}
	fmt.Println("jolts retruned as: ", result)
	return result
}

func findLargestJolt(jolts []int, number, rem int) int {
	totalElement := 12
	elementLeft := 11
	for idx := range totalElement {
		if number > jolts[idx] && rem >= elementLeft {
			jolts[idx] = number
			reset(jolts, idx+1)
		}
	}
}

func D3P1(inputFile string) {
	totOutputJolate := 0
	for bank := range strings.SplitSeq(inputFile, "\n") {
		tenthPlace := 0
		unitPlace := 0

		var largestJoleInBank int
		for i := range len(bank) {
			number, _ := strconv.Atoi(string(bank[i]))
			if number > tenthPlace &&
				i != len(bank)-1 { // so that 10th place will never be at the last of the bank
				tenthPlace = number
				unitPlace = 0 // reseting unitplace so that it doesn't carry anything that's preceeding current 1-th place

			} else if number > unitPlace {
				unitPlace = number
			}

			largestJoleInBank = tenthPlace*10 + unitPlace
		}

		totOutputJolate += largestJoleInBank
		// fmt.Printf("%v: %v\n", bank, largestJoleInBank)
	}

	fmt.Println("Day 3, P1: ", totOutputJolate)
}

func D3P2(inputFile string) {
	totOutputJolate12Dig := 0

	for bank := range strings.SplitSeq(inputFile, "\n") {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.")
		fmt.Println(bank)
		jolts := make([]int, 12)

		for i := range len(bank) {
			number, _ := strconv.Atoi(string(bank[i]))
			remaining := len(bank) - i - 1

			findLargestJolt(jolts, number, remaining)
			if number > jolts[0] && remaining >= 11 {
				jolts[0] = number
				reset(jolts, 1)
			} else if number > jolts[1] && remaining >= 10 {
				jolts[1] = number
				reset(jolts, 2)
			} else if number > jolts[2] && remaining >= 9 {
				jolts[2] = number
				reset(jolts, 3)
			} else if number > jolts[3] && remaining >= 8 {
				jolts[3] = number
				reset(jolts, 4)
			} else if number > jolts[4] && remaining >= 7 {
				jolts[4] = number
				reset(jolts, 5)
			} else if number > jolts[5] && remaining >= 6 {
				jolts[5] = number
				reset(jolts, 6)
			} else if number > jolts[6] && remaining >= 5 {
				jolts[6] = number
				reset(jolts, 7)
			} else if number > jolts[7] && remaining >= 4 {
				jolts[7] = number
				reset(jolts, 8)
			} else if number > jolts[8] && remaining >= 3 {
				jolts[8] = number
				reset(jolts, 9)
			} else if number > jolts[9] && remaining >= 2 {
				jolts[9] = number
				reset(jolts, 10)
			} else if number > jolts[10] && remaining >= 1 {
				jolts[10] = number
				reset(jolts, 11)
			} else if number > jolts[11] && remaining >= 0 {
				jolts[11] = number
				reset(jolts, 12)
			}

			fmt.Println("Jolts at end: ", jolts)

		}

		largestJoltInBank := formNumber(jolts)
		// fmt.Printf("%v\n%v\n\n", bank, largestJoltInBank)
		totOutputJolate12Dig += largestJoltInBank
		fmt.Println("tot: ", totOutputJolate12Dig)
	}

	fmt.Println("D3P2: ", totOutputJolate12Dig)
}

func main() {
	inputFile := strings.TrimSpace(utils.ReadFile("../inputs/day_3.txt"))
	D3P1(inputFile)
	D3P2(inputFile)
}
