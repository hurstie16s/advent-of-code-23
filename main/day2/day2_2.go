package day2

import (
	shared "advent-of-code-23/main/shared_code"
	"fmt"
)

func Call_2_2() {

	//filePath := "main/day2/day2_eg.txt"
	filePath := "main/day2/day2.txt"
	//filePath := "main/day2/day2_testing.txt"

	fileContents := shared.GetFileContents(filePath)

	games := getGameMaps(fileContents)

	powerGames := calculatePower(games)

	result := shared.Sum(powerGames)

	fmt.Printf("Power of games: %d\n", result)

}

func calculatePower(games map[int][]int) []int {

	var powers []int

	for _, values := range games {
		var gamePower = 1
		for _, val := range values {
			gamePower *= val
		}
		powers = append(powers, gamePower)
	}

	return powers
}
