package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func D1P1(inputFile string) {
	dialPos := 50
	totalZeroHits := 0

	for line := range strings.SplitSeq(inputFile, "\n") {
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

	fmt.Println("Day 1, Part 1: ", totalZeroHits)
}

func D1P2(inputFile string) {
	dialPos := 50
	totZeroVisted := 0
	oldDialPos := dialPos

	for line := range strings.SplitSeq(inputFile, "\n") {
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

	fmt.Println("Day 2, Part 2: ", totZeroVisted)
}
