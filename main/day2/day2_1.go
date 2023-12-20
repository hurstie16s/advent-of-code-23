package day2

import (
	shared "advent-of-code-23/main/shared_code"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
// Answer too high
51 is incorrectly accepted
*/

const MaxRed = 12
const MaxGreen = 13
const MaxBlue = 14

func Call_2_1() {

	//filePath := "main/day2/day2_eg.txt"
	filePath := "main/day2/day2.txt"
	//filePath := "main/day2/day2_testing.txt"

	fileContents := shared.GetFileContents(filePath)

	games := getGameMaps(fileContents)

	possibleGames := getGamesPossible(games)
	sort.Ints(possibleGames)
	fmt.Printf("Possible Game Ids: %v\n", possibleGames)

	summedGames := shared.Sum(possibleGames)

	fmt.Printf("Sum of possible games: %d\n", summedGames)

}

func getColourCount(sets []string) (int, int, int) {

	var blueCount int
	var redCount int
	var greenCount int

	for _, set := range sets {
		splitSets := strings.Split(set, ", ")
		for _, colour := range splitSets {
			colourSplit := strings.Split(colour, " ")
			val, _ := strconv.Atoi(colourSplit[0])
			switch colourSplit[1] {
			case "blue":
				blueCount = int(math.Max(float64(blueCount), float64(val)))
			case "red":
				redCount = int(math.Max(float64(redCount), float64(val)))
			case "green":
				greenCount = int(math.Max(float64(greenCount), float64(val)))
			}
		}
	}

	return blueCount, redCount, greenCount
}

func getGameMaps(gameStrings []string) map[int][]int {

	var games = make(map[int][]int)

	for i := 0; i < len(gameStrings); i++ {
		splitLine := strings.Split(gameStrings[i], ": ")
		id, _ := strconv.Atoi(strings.Split(splitLine[0], " ")[1])
		sets := strings.Split(splitLine[1], "; ")

		blueCount, redCount, greenCount := getColourCount(sets)

		gameValues := []int{blueCount, redCount, greenCount}

		// Map of game to sets
		games[id] = gameValues
	}

	return games
}

func getGamesPossible(games map[int][]int) []int {
	var possibleIds []int

	for id, values := range games {
		fmt.Printf("Id: %d, Blue: %d, Red: %d, Green %d\n", id, values[0], values[1], values[2])
		if values[0] <= MaxBlue &&
			values[1] <= MaxRed &&
			values[2] <= MaxGreen {
			possibleIds = append(possibleIds, id)
		}
	}

	return possibleIds
}
