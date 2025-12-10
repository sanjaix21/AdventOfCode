package days

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseOp(line string) []string {
	line = strings.TrimSpace(line)
	parts := strings.Fields(line)

	out := make([]string, len(parts))
	for i, p := range parts {
		out[i] = string(p)
	}

	return out
}

func parseNumPadding(line string) []string {
	line = strings.TrimSpace(line)
	parts := strings.Fields(line)

	out := make([]string, len(parts))
	for i, p := range parts {
		n, _ := strconv.Atoi(p)
		out[i] = fmt.Sprintf("%04d", n)
	}

	return out
}

func parseInts(line string) []int {
	line = strings.TrimSpace(line)
	parts := strings.Fields(line)

	out := make([]int, len(parts))

	for i, p := range parts {
		out[i], _ = strconv.Atoi(p)
	}

	return out
}

func D6P1(inputFile string) {
	lines := slices.Collect(strings.SplitSeq(inputFile, "\n"))
	opLine := parseOp(lines[4])
	totalOps := len(opLine)
	grandTotal := 0

	for i := range totalOps {
		op := opLine[i]
		tempAdd := 0
		tempMul := 1
		for j := range 4 {
			if op == "+" {
				tempAdd += parseInts(lines[j])[i]
			} else if op == "*" {
				tempMul *= parseInts(lines[j])[i]
			}
		}

		if op == "+" {
			grandTotal += tempAdd
		} else if op == "*" {
			grandTotal += tempMul
		}
	}

	fmt.Println("Day 6, Part 1: ", grandTotal)
}
