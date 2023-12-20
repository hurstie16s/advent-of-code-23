package day2

import (
	shared "advent-of-code-23/main/shared_code"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const MaxRed = 12
const MaxGreen = 13
const MaxBlue = 14

func Call_2_1() {

	filePath := "main/day2/day2_eg.txt"

	fileContents := shared.GetFileContents(filePath)

	games := getGameMaps(fileContents)

	possibleGames := getGamesPossible(games)

	summedGames := shared.Sum(possibleGames)

	fmt.Printf("Sum of possible games: %d\n", summedGames)

}

func getColourCount(colour string, sets []string) int {

	var count int

	for _, set := range sets {
		if strings.Contains(set, colour) {
			index := strings.Index(set, colour) - 2 // TODO: Fix this so it handles 2 digit numbers
			val, _ := strconv.Atoi(strings.TrimSpace(set[index : index+1]))
			count = int(math.Max(float64(count), float64(val)))
		}
	}

	return count
}

func getGameMaps(gameStrings []string) map[int][]int {

	var games = make(map[int][]int)

	for i := 0; i < len(gameStrings); i++ {
		splitLine := strings.Split(gameStrings[i], ": ")
		id, _ := strconv.Atoi(strings.Split(splitLine[0], " ")[1])
		sets := strings.Split(splitLine[1], ";")

		blueCount := getColourCount("blue", sets)
		redCount := getColourCount("red", sets)
		greenCount := getColourCount("green", sets)

		gameValues := []int{blueCount, redCount, greenCount}

		// Map of game to sets
		games[id] = gameValues
	}

	return games
}

func getGamesPossible(games map[int][]int) []int {
	var possibleIds []int

	for id, values := range games {
		fmt.Printf("ID: %d, Values: %v\n", id, values)
		if values[0] <= MaxBlue &&
			values[1] <= MaxRed &&
			values[2] <= MaxGreen {
			possibleIds = append(possibleIds, id)
			fmt.Println(id)
		}
	}

	return possibleIds
}
