package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Part01 answer: 11567
func Part01(content string) int {
	instructions, contentMap := parseContent(content)

	found := false
	i := 0
	next := "AAA"
	target := "ZZZ"
	for !found {
		for _, ins := range instructions {
			i++
			switch ins {
			case "L":
				next = contentMap[next][0]
			case "R":
				next = contentMap[next][1]
			}

			if next == target {
				found = true
			}
		}
	}

	return i
}

// Part02 answer: 9858474970153
func Part02(content string) int {
	instructions, contentMap := parseContent(content)

	var startingNodes []string
	for key := range contentMap {
		if string(key[len(key)-1]) == "A" {
			startingNodes = append(startingNodes, key)
		}
	}

	count := len(instructions)
	for _, node := range startingNodes {
		var cycles int
		for string(node[len(node)-1]) != "Z" {
			for _, ins := range instructions {
				switch ins {
				case "L":
					node = contentMap[node][0]
				case "R":
					node = contentMap[node][1]
				}
			}
			cycles++
		}
		count *= cycles
	}

	return count
}

func parseContent(content string) ([]string, map[string][]string) {
	blocks := strings.Split(content, "\n\n")
	instructions := strings.Split(blocks[0], "")

	lines := strings.Split(blocks[1], "\n")
	contentMap := make(map[string][]string, len(lines))

	for _, line := range lines {
		key, parts := parseLine(line)
		contentMap[key] = parts
	}

	return instructions, contentMap
}

func parseLine(line string) (string, []string) {
	parts := strings.Split(line, "=")

	key := strings.TrimSpace(parts[0])

	next := parts[len(parts)-1]
	next = strings.Replace(next, "(", "", 1)
	next = strings.Replace(next, ")", "", 1)
	next = strings.ReplaceAll(next, " ", "")

	return key, strings.Split(next, ",")
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
