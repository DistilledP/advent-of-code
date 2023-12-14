package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Part01: 861300
func Part01(content string) int {
	raceHistory := parseInput(content)
	var winningCombinations []int

	for _, race := range raceHistory {
		held := 0
		winners := 0
		for held < race.time {
			distanceTravelled := calculateDistanceTravelled(held, race.time)
			if distanceTravelled > race.distance {
				winners++
			}
			held++
		}

		winningCombinations = append(winningCombinations, winners)
	}

	total := 1
	for _, combination := range winningCombinations {
		total *= combination
	}

	return total
}

func Part02(content string) int {

	return 0
}

type timeDistance struct {
	time     int
	distance int
}

func parseInput(content string) []timeDistance {
	content = strings.ReplaceAll(content, "\t", ",")
	lines := strings.Split(content, "\n")

	var rawContent [][]int = make([][]int, 2)
	for i := 0; i < len(lines); i++ {
		col := 0
		parts := strings.Split(lines[i], " ")
		for j := 0; j < len(parts); j++ {
			if num, err := strconv.Atoi(parts[j]); err == nil {
				rawContent[i] = append(rawContent[i], num)
				col++
			}
		}
	}

	var timeDistances []timeDistance
	for idx := range rawContent[0] {
		timeDistances = append(timeDistances, timeDistance{
			time:     rawContent[0][idx],
			distance: rawContent[1][idx],
		})
	}

	return timeDistances
}

func calculateDistanceTravelled(msHeld, raceLength int) int {
	return msHeld * (raceLength - msHeld)
}

// ================================================
// Helper fuctions
// ================================================
func dumpAsJson(data interface{}, filename string) {
	dataJson, _ := json.MarshalIndent(data, "", "  ")

	cwd, _ := os.Getwd()
	err := os.WriteFile(fmt.Sprintf("%s/../../out/%s.json", cwd, filename), dataJson, 0644)
	if err != nil {
		log.Fatalf("failed to write file: %s", err)
	}
}

func debugLog(items ...any) {
	if os.Getenv("DEBUG") == "1" {
		log.Println(items...)
	}
}

func mustInt(target string) int {
	if num, err := strconv.Atoi(target); err == nil {
		return num
	}

	return 0
}

var spinnerFrames = []string{
	"/", "─", "\\", "│",
}
var spinnerFrame int

func spinnerNextFrame() {
	if spinnerFrame >= len(spinnerFrames) {
		spinnerFrame = 0
	}

	fmt.Printf("\033[1m\033[7m\r%s\r\033[0m", spinnerFrames[spinnerFrame])

	spinnerFrame++
}
