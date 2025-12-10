package days

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func D5P2(inputFile string) {
	// extract the first block (fresh ranges section)
	block := slices.Collect(strings.SplitSeq(inputFile, "\n\n"))[0]
	lines := slices.Collect(strings.SplitSeq(block, "\n"))

	freshRanges := [][]int{}

	// parse "m-n"
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := slices.Collect(strings.SplitSeq(line, "-"))
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		freshRanges = append(freshRanges, []int{start, end})
	}

	// sort by start
	slices.SortFunc(freshRanges, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	// merge and count
	total := 0
	curStart := freshRanges[0][0]
	curEnd := freshRanges[0][1]

	for i := 1; i < len(freshRanges); i++ {
		r := freshRanges[i]

		if r[0] <= curEnd+1 { // overlap or touching
			if r[1] > curEnd {
				curEnd = r[1]
			}
		} else {
			// finalize previous range
			total += (curEnd - curStart + 1)
			curStart = r[0]
			curEnd = r[1]
		}
	}

	// add last range
	total += (curEnd - curStart + 1)

	fmt.Println(total)
}

func D5P1(inputFile string) {
	parts := slices.Collect(strings.SplitSeq(inputFile, "\n\n"))
	freshRanges := parts[0]
	ingdxIds := parts[1]
	totFreshIngdx := 0

	for ingdxId := range strings.SplitSeq(ingdxIds, "\n") {
		ingdxId, _ := strconv.Atoi(ingdxId)
		for vRange := range strings.SplitSeq(freshRanges, "\n") {
			parts := slices.Collect(strings.SplitSeq(string(vRange), "-"))
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			if ingdxId >= start && ingdxId <= end {
				totFreshIngdx++
				break
			}
		}
	}

	fmt.Println(totFreshIngdx)
}
