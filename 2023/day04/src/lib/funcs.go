package lib

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	ID       int
	Winners  []int
	Selected []int
	Count    int
}

// Part01 answer: 28538
func Part01(content string) int {
	cards := parseData(content)

	runningTotal := 0
	for _, card := range cards {
		winningNumbers := intersect(card.Winners, card.Selected)

		lineTotal := 0
		if len(winningNumbers) > 0 {
			for i := range winningNumbers {
				if i < 2 {
					lineTotal += 1
				} else {
					// shift the bits of the number 2 to the left by (i - 2) positions.
					lineTotal += 2 << (i - 2)
				}
			}
		}

		runningTotal += lineTotal

		debugOutput := fmt.Sprintf(
			`{ID: %d, Winners: %v, Selected: %v, Intersect: %v, lineTotal: %d, runningTotal: %d}`,
			card.ID, card.Winners, card.Selected, winningNumbers, lineTotal, runningTotal,
		)

		debugLog(debugOutput)
	}

	return runningTotal
}

// Part02 answer: 9425061
func Part02(content string) int {
	cards := parseData(content)
	runningTotal := 0

	var cardCount map[int]int = make(map[int]int, 1000)
	for cID := 1; cID <= len(cards); cID++ {
		card := cards[cID]

		cardCount[card.ID]++

		winningNumbers := intersect(card.Winners, card.Selected)
		for j := 0; j < len(winningNumbers); j++ {
			newIdx := card.ID + j + 1
			cardCount[newIdx] += cardCount[card.ID]
		}
	}

	for _, c := range cardCount {
		runningTotal += c
	}

	return runningTotal
}

func intersect(set1, set2 []int) []int {
	var common []int

	for _, i := range set1 {
		if slices.Contains(set2, i) {
			common = append(common, i)
		}
	}

	return common
}

func parseData(content string) map[int]Card {
	var cards map[int]Card = make(map[int]Card)

	for _, line := range strings.Split(content, "\n") {
		newCard := Card{}

		card := strings.Split(line, ":")

		cardParts := strings.Split(card[0], " ")
		slices.Reverse(cardParts)

		if cardId, err := strconv.Atoi(cardParts[0]); err == nil {
			newCard.ID = cardId
		}

		numbers := strings.Split(card[1], "|")

		newCard.Winners = extractNumbers(numbers[0])
		newCard.Selected = extractNumbers(numbers[1])
		newCard.Count = len(intersect(newCard.Winners, newCard.Selected))

		cards[newCard.ID] = newCard
	}

	return cards
}

func extractNumbers(numbers string) []int {
	var nums []int
	for _, number := range strings.Split(numbers, " ") {
		if num, err := strconv.Atoi(number); err == nil {
			nums = append(nums, num)
		}
	}

	return nums
}

func debugLog(items ...any) {
	if os.Getenv("DEBUG") == "1" {
		fmt.Println(items...)
	}
}
