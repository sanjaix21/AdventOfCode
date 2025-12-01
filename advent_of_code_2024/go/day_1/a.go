package main

import (
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() string {
	data, err := os.ReadFile("input.txt")
}

func main() {
	data, err := os.ReadFile("./input.txt")
	checkErr(err)
	data = data.Tr
	fmt.Println(string(data))

}
