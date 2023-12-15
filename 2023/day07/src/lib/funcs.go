package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Part01 answer: 250474325
func Part01(content string) int {
	games := parseContent(content, false)
	slices.SortFunc[[]hand](games, cardSort)

	return calculateWinnings(games)
}

// Part02 answer:
func Part02(content string) int {
	cardValues["J"] = 1
	defer func() {
		cardValues["J"] = 11
	}()

	games := parseContent(content, true)
	slices.SortFunc[[]hand](games, cardSort)

	return calculateWinnings(games)
}

var cardValues map[string]uint = make(map[string]uint, 15)

func init() {
	cardValues = map[string]uint{
		"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9,
		"8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
	}
}

type hand struct {
	cards      string
	bid        int
	cardsValue uint
}

func parseContent(content string, part2 bool) []hand {
	var hands []hand
	lines := strings.Split(content, "\n")

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		hands = append(hands, hand{
			cards:      parts[0],
			bid:        mustInt(parts[1]),
			cardsValue: calculateHandValue(parts[0], part2),
		})
	}

	return hands
}

func calculateHandValue(hand string, part2 bool) uint {
	freqs, max, jacks := cardFreq(hand, part2)

	freqsLen := uint(len(freqs))
	if freqsLen == 0 {
		freqsLen = 1
	}

	return (((max+jacks)<<3)-freqsLen)<<20 |
		cardValues[string(hand[0])]<<16 |
		cardValues[string(hand[1])]<<12 |
		cardValues[string(hand[2])]<<8 |
		cardValues[string(hand[3])]<<4 |
		cardValues[string(hand[4])]
}

func cardFreq(hand string, sumJacks bool) (map[string]uint, uint, uint) {
	cardFreqMap := make(map[string]uint, 13)
	var max, jacks uint

	for _, card := range hand {
		cardStr := string(card)
		if sumJacks && cardStr == "J" {
			jacks++
		} else {
			cardFreqMap[cardStr]++
			if cardFreqMap[cardStr] > max {
				max = cardFreqMap[cardStr]
			}
		}
	}

	return cardFreqMap, max, jacks
}

func calculateWinnings(games []hand) int {
	total := 0

	for i, game := range games {
		total += game.bid * (i + 1)
	}

	return total
}

func cardSort(a, b hand) int {
	if a.cardsValue == b.cardsValue {
		return 0
	}

	if a.cardsValue > b.cardsValue {
		return 1
	}

	return -1
}

// ================================================
// Helper fuctions
// ================================================

func mustInt(val string) int {
	if num, err := strconv.Atoi(val); err == nil {
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
