package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := readFromFile("./test_input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println(content)
	for _, line := range content {
		// numbers are bytes 48-57 (0-9)
	}
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
