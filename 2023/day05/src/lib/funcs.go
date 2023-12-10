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

/**

Current parsing/logic+understanding is completely wrong.

It needs to be a map from one to the other, for example:

seed-to-soil:
98 -> 50
99 -> 51
50 -> 52
51 -> 53
52 -> 54
...

soil-to-fertilizer:
15 -> 0
16 -> 1
17 -> 2
...
52 -> 37
53 -> 38
0  -> 39
1  -> 40
...

fertilizer-to-water:
53 -> 49
54 -> 50
...
11 -> 0
12 -> 1
...
0 -> 42
1 -> 43
...
7 -> 57
8 -> 58
...

etc etc.

If there isn't a mapping then source + destination are the same, effectively the mapping is an override.

Might as well start again :D

Update: current solution should work, however due the amount of numbers involved it's slow.  Need to look at mapping ranges rather than numbers.

Part01 done... now for part 2
*/

// Part01 answer: 324724204
func Part01(content string) int {
	seeds, mapping2 := buildMap(content)

	var locations []int
	for i := 0; i < len(seeds); i++ {
		needle := seeds[i]
		for j := 0; j < len(mapping2); j++ {
			resp := walkMap(mapping2, j, needle)
			needle = resp[0]
		}

		locations = append(locations, needle)
	}

	return slices.Min(locations)
}

func Part02(content string) int {

	return 0
}

type Mapping struct {
	Start       map[int]int
	End         map[int]int
	SourceStart int
	SourceEnd   int
	Length      int
}

func walkMap(builtMap [][]Mapping, idx, needle int) []int {
	var response []int

	mapping := builtMap[idx]

	for i := 0; i < len(mapping); i++ {
		rangeMap := mapping[i]

		if needle >= rangeMap.SourceStart && needle <= rangeMap.SourceEnd {
			offset := needle - rangeMap.SourceStart
			return []int{rangeMap.Start[rangeMap.SourceStart] + offset}
		}
	}

	response = append(response, needle)

	return response
}

func buildMap(content string) ([]int, [][]Mapping) {
	blocks := strings.Split(content, "\n\n")

	// var output [][]Mapping
	output := make([][]Mapping, len(blocks))
	for i := 0; i < len(blocks); i++ {
		blockLines := strings.Split(blocks[i], "\n")

		// debugLog(blockLines)

		blockMapping := make([]Mapping, 0)
		for j := 1; j < len(blockLines); j++ {
			lineParts := strings.Split(blockLines[j], " ")

			src := MustInt(lineParts[1])
			dst := MustInt(lineParts[0])
			rangeLength := MustInt(lineParts[2])

			blockMapping = append(blockMapping, Mapping{
				Start:       map[int]int{src: dst},
				End:         map[int]int{src + rangeLength - 1: dst + rangeLength - 1},
				SourceStart: src,
				SourceEnd:   src + rangeLength - 1,
				Length:      rangeLength,
			})
		}

		output[i] = blockMapping
	}

	if len(output[0]) == 0 {
		output = output[1:]
	}
	return processSeedBlock(blocks[0]), output
}

func processSeedBlock(block string) []int {
	cleanBlock := strings.TrimPrefix(block, "seeds:")

	seeds := []int{}
	for _, seed := range strings.Split(cleanBlock, " ") {
		if num, err := strconv.Atoi(seed); err == nil {
			seeds = append(seeds, num)
		}
	}

	return seeds
}

// ================================================
// Helper fuctions
// ================================================
func MustInt(target string) int {
	if num, err := strconv.Atoi(target); err == nil {
		return num
	}

	return -1
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
