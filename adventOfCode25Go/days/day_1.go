package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadFile() string {
	filePath := "../inputs/day_1.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func Part1() {
	inputData := ReadFile()
	dialPos := 50
	totalZeroHits := 0

	for line := range strings.SplitSeq(inputData, "\n") {
		if len(line) > 0 {
			direction := line[:1]
			magnitude, _ := strconv.Atoi(line[1:])

			if direction == "R" {
				dialPos = (dialPos + magnitude) % 100
			} else if direction == "L" {
				dialPos = (dialPos - magnitude) % 100
			}

			if dialPos == 0 {
				totalZeroHits++
			}
		}
	}

	fmt.Println("Par1 Ans: ", totalZeroHits)
}

func Part2() {
	inputData := ReadFile()
	dialPos := 50
	totZeroVisted := 0
	oldDialPos := dialPos

	for line := range strings.SplitSeq(inputData, "\n") {
		if len(line) > 0 {
			direction := line[:1]
			magnitude, _ := strconv.Atoi(line[1:])

			if direction == "R" {
				dialPos += magnitude
			} else if direction == "L" {
				dialPos -= magnitude
			}

			if dialPos > 99 {
				totZeroVisted += int(math.Abs(float64(dialPos)) / 100)
			}

			if dialPos < 1 {
				totZeroVisted += int(math.Abs(float64(dialPos))/100) + 1
				if oldDialPos == 0 {
					totZeroVisted -= 1
				}
			}

			dialPos = (dialPos%100 + 100) % 100
			oldDialPos = dialPos
		}
	}

	fmt.Printf("Part2 Ans: %v\n", totZeroVisted)
}
