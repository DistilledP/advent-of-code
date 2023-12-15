package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 1984151189 too high
func Part01(content string) int64 {
	lines := parseContent(content)

	var sum int64
	for _, line := range lines {
		var lineDeltas [][]int64
		lineDeltas = append(lineDeltas, line)
		delta := deltas(line)
		lineDeltas = append(lineDeltas, delta)
		for !allZeros(delta) {
			delta = deltas(delta)
			lineDeltas = append(lineDeltas, delta)
		}
		sum += nextValue(lineDeltas)
	}

	return sum
}

func Part02(content string) int {

	return 0
}

func parseContent(content string) [][]int64 {
	lines := strings.Split(content, "\n")

	var numLines [][]int64
	for _, line := range lines {

		parts := strings.Split(line, " ")
		var nums []int64
		for i := 0; i < len(parts); i++ {
			if num, err := strconv.ParseInt(string(parts[i]), 10, 64); err == nil {
				nums = append(nums, num)
			}
		}

		numLines = append(numLines, nums)
	}

	return numLines
}

func deltas(sub []int64) []int64 {
	var deletas []int64
	for i := 0; i < len(sub); i++ {
		lower := i
		upper := i + 1
		if upper > len(sub)-1 {
			break
		}

		deletas = append(deletas, diff(sub[lower], sub[upper]))
	}

	return deletas
}

func diff(a, b int64) int64 {
	if a < b {
		return b - a
	}
	return a - b
}

func allZeros(sub []int64) bool {
	var sum int64
	for _, num := range sub {
		sum += num
	}

	return sum == 0
}

func nextValue(deltas [][]int64) int64 {

	// debugLog("----------------------------")
	// debugLog("start", deltas)

	var nextVal int64
	for i := len(deltas) - 2; i >= 0; i-- {
		// debugLog(deltas[i])

		if i-1 >= 0 {
			offset := deltas[i][len(deltas[i])-1]
			preVal := deltas[i-1][len(deltas[i])]

			nextVal = preVal + offset

			if i-1 > 0 {
				deltas[i-1] = append(deltas[i], nextVal)
				// offset = nextVal
			}

			// debugLog("offset", offset, "preVal", preVal, "nextVal", nextVal, deltas)
		}

		// nextVal = deltas[i-1][[len(deltas[i])-1]] + deltas[i][len(deletas[i])-1]
	}

	// debugLog("finish", deltas)

	return nextVal
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
