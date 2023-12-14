package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Part01(content string) int {
	games := parseContent(content)

	// Sorting of the games is incorrect when using the full data set (it's the only thing that makes sense).
	// Need to examine the hands in more detail to correct the sorting.
	// Maybe build up a map of doubles, triples etc for each hand and use that to determine the order.
	// Scoring the entire hand alone does not do what we need it to do (although the test passes).
	slices.SortFunc[[]hand](games, func(a, b hand) int {
		cardCountsA := countsOfCards(a.cards)
		cardCountsB := countsOfCards(b.cards)

		debugLog(cardCountsA, cardCountsB)

		return a.cardsValue - b.cardsValue
	})

	debugLog(games)

	return calculateWinnings(games)
}

func Part02(content string) int {

	return 0
}

var cardValues map[string]int = make(map[string]int, 15)

func init() {
	cardValues = map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}
}

type hand struct {
	cards      string
	bid        int
	cardsValue int
}

func parseContent(content string) []hand {
	var hands []hand
	lines := strings.Split(content, "\n")

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		hands = append(hands, hand{
			cards:      parts[0],
			bid:        mustInt(parts[1]),
			cardsValue: calculateHandValue(parts[0]),
		})
	}

	return hands
}

func calculateHandValue(hand string) int {
	cardCounts := countsOfCards(hand)
	cardsVal := 0
	pos := len(hand)
	for j := 0; j < len(hand); j++ {
		card := string(hand[j])

		cardVal := cardValues[card]
		cardsVal += cardVal * cardCounts[card] * pos
		pos--
	}

	return cardsVal
}

func countsOfCards(hand string) map[string]int {
	var result map[string]int = make(map[string]int, 14)
	for c := range cardValues {
		regexp1 := regexp.MustCompile(fmt.Sprintf("[%s]", c))
		cardCount := len(regexp1.FindAllString(hand, -1))

		var cardMulti int
		switch cardCount {
		case 2:
			cardMulti = 4
		case 3:
			cardMulti = 16
		case 4:
			cardMulti = 64
		case 5:
			cardMulti = 128
		default:
			cardMulti = 1
		}
		result[c] = cardCount * cardMulti
	}

	return result
}

func calculateWinnings(games []hand) int {
	var total int

	for i, game := range games {
		total += game.bid * (i + 1)
	}

	return total
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
