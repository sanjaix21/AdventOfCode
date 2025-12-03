package utils

import "os"

func ReadFile(inputFilePath string) string {
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func ReverseString(word string) string {
	runes := []rune(word)
	var reversed string

	for i := len(runes) - 1; i >= 0; i-- {
		reversed += string(runes[i])
	}

	return reversed
}
