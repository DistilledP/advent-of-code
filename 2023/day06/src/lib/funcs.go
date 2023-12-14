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
func Part01(content string) uint64 {
	return calculateWinners(parseInput(content))
}

// Part02: 28101347
func Part02(content string) uint64 {
	return calculateWinners(parseInputPart2(content))
}

type timeDistance struct {
	time     uint64
	distance uint64
}

func parseInput(content string) []timeDistance {
	lines := strings.Split(content, "\n")

	var rawContent [][]uint64 = make([][]uint64, 2)
	for i := 0; i < len(lines); i++ {
		col := 0
		parts := strings.Split(lines[i], " ")
		for j := 0; j < len(parts); j++ {
			if num, err := strconv.ParseUint(parts[j], 10, 64); err == nil {
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

func parseInputPart2(content string) []timeDistance {
	lines := strings.Split(content, "\n")

	var timeDistances []timeDistance
	timeDistances = append(timeDistances, timeDistance{
		time:     concatLine(lines[0]),
		distance: concatLine(lines[1]),
	})

	return timeDistances
}

func calculateDistanceTravelled(msHeld, raceLength uint64) uint64 {
	return msHeld * (raceLength - msHeld)
}

func calculateWinners(raceHistory []timeDistance) uint64 {
	var winningCombinations []uint64
	for _, race := range raceHistory {
		var held uint64 = 0
		var winners uint64 = 0
		for held < race.time {
			distanceTravelled := calculateDistanceTravelled(held, race.time)
			if distanceTravelled > race.distance {
				winners++
			}
			held++
		}

		winningCombinations = append(winningCombinations, winners)
	}

	var total uint64 = 1
	for _, combination := range winningCombinations {
		total *= combination
	}

	return total
}

// ================================================
// Helper fuctions
// ================================================

func concatLine(theLine string) uint64 {
	line := strings.Split(theLine, " ")
	numbers := strings.TrimSpace(strings.ReplaceAll(strings.Join(line[1:], ""), " ", ""))

	return mustUInt64(numbers)
}

func mustUInt64(val string) uint64 {
	if num, err := strconv.ParseUint(val, 10, 64); err == nil {
		return num
	}

	return 0
}

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
