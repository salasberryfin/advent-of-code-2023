package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getCalibrationPart1(filename string) int {
	content, err := readFromFile(filename)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var calibration int
	for _, line := range content {
		// numbers are bytes 49-57 (1-9)
		var numbers []int
		for i := 0; i < len(line); i++ {
			if line[i] >= 49 && line[i] <= 57 {
				numbers = append(numbers, int(line[i])-48)
			}
		}
		calibration += numbers[0]*10 + numbers[len(numbers)-1]
	}

	fmt.Println("Calibration: ", calibration)

	return calibration
}

func getCalibrationPart2(filename string) int {
	content, err := readFromFile(filename)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	wordsMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var calibration int

	for _, line := range content {
		fmt.Println("line:", line)
		for digit, word := range wordsMap {
			line = strings.Replace(line, digit, word, -1)
		}
		fmt.Println("new line:", line)
	}

	fmt.Println("Calibration: ", calibration)

	return calibration
}

func readFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	content := []string{}
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content, nil
}
