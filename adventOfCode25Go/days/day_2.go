package main

import (
	"fmt"
	"strconv"
	"strings"
)

func D2P1(inputFile string) {
	idRanges := strings.Split(inputFile, ",")
	totInvalidIdVal := 0

	for _, ids := range idRanges {
		parts := strings.Split(ids, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for id := start; id <= end; id++ {
			idStr := strconv.Itoa(id)
			if len(idStr)%2 == 0 {
				mid := len(idStr) / 2
				if idStr[:mid] == idStr[mid:] {
					totInvalidIdVal += id
				}
			}
		}
	}

	fmt.Println("Part1: ", totInvalidIdVal)
}

func D2P2(inputFile string) {
	idRanges := strings.Split(inputFile, ",")
	totInvalidIdVal := 0
	for _, ids := range idRanges {
		parts := strings.Split(ids, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for id := start; id <= end; id++ {
			idStr := strconv.Itoa(id)
			strLen := len(idStr)
			for window := 1; window <= strLen/2; window++ {
				if strLen%window != 0 {
					continue
				}
				subString := idStr[:window]
				repeats := strLen / window
				if strings.Repeat(subString, repeats) == idStr {
					fmt.Printf(
						"SubsStr: %v (%v), len: %v, wind: %v, :[%v]\n",
						subString,
						repeats,
						len(idStr),
						window+1,
						id,
					)
					totInvalidIdVal += id
					break
				}
			}
		}
	}

	fmt.Println("Part 2: ", totInvalidIdVal)
}
