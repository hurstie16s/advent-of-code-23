package day1

// System Imports
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Call_1_2() {
	var result int

	numbersAsStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numbersAsValues := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	filePath := "main/day1/task1_1.txt"
	//filePath := "main/day1/task1_2_eg.txt"

	fileContents := getFileContents(filePath)
	firstNums := parseForFirst(numbersAsStrings, numbersAsValues, fileContents)
	lastNums := parseForLast(numbersAsStrings, numbersAsValues, fileContents)
	nums := zipNums(firstNums, lastNums)
	result = sum(nums)
	fmt.Printf("Result: %d", result)
}

func getFileContents(path string) []string {
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

	var input []string
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}

	return input
}

func parseForFirst(
	numsAsStrings []string,
	numsAsValues []string,
	fileContents []string) []string {

	var firstNums []string

	for i := 0; i < len(fileContents); i++ {

		var firstNum string
		var firstNumIndex = -1

		line := fileContents[i]

		for j := 0; j < len(numsAsStrings); j++ {
			numString := numsAsStrings[j]
			if strings.Contains(line, numString) {
				index := strings.Index(line, numString)
				if firstNumIndex == -1 || index < firstNumIndex {
					firstNumIndex = index
					firstNum = numsAsValues[j]
				}
			}
		}

		for j := 0; j < len(numsAsStrings); j++ {
			numVal := numsAsValues[j]
			if strings.Contains(line, numVal) {
				index := strings.Index(line, numVal)
				if firstNumIndex == -1 || index < firstNumIndex {
					firstNumIndex = index
					firstNum = numsAsValues[j]
				}
			}
		}
		firstNums = append(firstNums, firstNum)
	}

	return firstNums
}

func parseForLast(
	numsAsStrings []string,
	numsAsValues []string,
	fileContents []string) []string {

	var lastNums []string

	for i := 0; i < len(fileContents); i++ {

		var lastNum string
		var lastNumIndex = -1

		line := fileContents[i]

		for j := 0; j < len(numsAsStrings); j++ {
			numString := numsAsStrings[j]
			if strings.Contains(line, numString) {
				index := strings.LastIndex(line, numString)
				if index > lastNumIndex {
					lastNumIndex = index
					lastNum = numsAsValues[j]
				}
			}
		}

		for j := 0; j < len(numsAsStrings); j++ {
			numVal := numsAsValues[j]
			if strings.Contains(line, numVal) {
				index := strings.LastIndex(line, numVal)
				if lastNumIndex == -1 || index > lastNumIndex {
					lastNumIndex = index
					lastNum = numsAsValues[j]
				}
			}
		}
		lastNums = append(lastNums, lastNum)
	}

	return lastNums
}

func zipNums(firstNums []string, lastNums []string) []int {
	var nums []int
	for i := 0; i < len(firstNums); i++ {
		newNum, _ := strconv.Atoi(firstNums[i] + lastNums[i])
		nums = append(nums, newNum)
		fmt.Printf("Num: %d\n", newNum)
	}
	return nums
}
