package shared_code

import (
	"bufio"
	"fmt"
	"os"
)

func GetFileContents(path string) []string {
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

func PrintFileContents(contents []string) {
	for _, line := range contents {
		fmt.Println(line)
	}
}
