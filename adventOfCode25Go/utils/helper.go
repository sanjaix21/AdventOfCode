package utils

import "os"

func ReadFile(inputFilePath string) string {
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	return string(data)
}
