package day1

// System Imports
import (
	shared "advent-of-code-23/main/shared_code"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Main function

func Call_1_1() {
	filePath := "task1_1.txt"

	finalNum := shared.Sum(getFileContentsProcessed(filePath))
	fmt.Printf("Final sum = %d", finalNum)
}

func getFileContentsProcessed(path string) []int {
	readFile, err := os.Open(path)
	defer func(readFile *os.File) {
		err := readFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(readFile)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var input []int

	for fileScanner.Scan() {
		line := strip(fileScanner.Text())
		firstNum := string(line[0])
		lastNum := string(line[len(line)-1])
		num, _ := strconv.Atoi(firstNum + lastNum)
		input = append(input, num)
	}

	return input
}

func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if '0' <= b && b <= '9' {
			result.WriteByte(b)
		}
	}
	return result.String()
}
