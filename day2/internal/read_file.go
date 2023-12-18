package internal

import (
	"bufio"
	"log"
	"os"
)

func ReadFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening the file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading lines", err)
	}

	return lines
}
