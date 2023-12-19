package main

// System Imports
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Main function
func main() {
	filePath := "task1_1.txt"
	readFile, err := os.Open(filePath)

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

	finalNum := sum(input)
	fmt.Printf("Final sum = %d", finalNum)
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

func sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
