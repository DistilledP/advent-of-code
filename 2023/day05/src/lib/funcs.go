package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
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

Part01 done.
Part02 done.
Part02 with batching done.
*/

// Part01 answer: 324724204
func Part01(content string) uint64 {
	seeds, mapping2 := buildMap(content)

	var lowest uint64 = math.MaxUint64
	for i := 0; i < len(seeds); i++ {
		needle := seeds[i]
		for j := 0; j < len(mapping2); j++ {
			resp := walkMap(mapping2, j, needle)
			needle = resp[0]
		}

		if needle < lowest {
			lowest = needle
		}
	}

	return lowest
}

// Part02 answer: 104070862
func Part02(content string) uint64 {
	seeds, mapping := buildMap(content)
	numberOfSeeds := len(seeds) / 2

	// brute force it.
	var lowest uint64 = math.MaxUint64
	for i := 0; i < len(seeds); i = i + 2 {
		currentSeed := int(math.Round(float64((i+1)/2))) + 1
		log.Printf("seed: %d/%d\tstart: %d\trange: %d\n", currentSeed, numberOfSeeds, seeds[i], seeds[i+1])
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			needle := j
			for k := 0; k < len(mapping); k++ {
				resp := walkMap(mapping, k, needle)
				needle = resp[0]
			}

			if needle < lowest {
				lowest = needle
			}
		}
	}

	return lowest
}

// Explore go routines option, synch.Waitgroup to batch the seeds in order
// to increase performance.
// Need to benchmark vs sequential.
func Part02_batch(content string) uint64 {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	seeds, mapping := buildMap(content)
	numberOfSeeds := len(seeds) / 2

	var seedLowest []uint64
	var batchSize int = 1000

	log.Println("batch size", batchSize)

	wg := sync.WaitGroup{}
	for i := 0; i < len(seeds); i = i + 2 {
		currentSeed := int(math.Round(float64((i+1)/2))) + 1
		log.Printf("seed: %d/%d\tstart: %d\trange: %d\n", currentSeed, numberOfSeeds, seeds[i], seeds[i+1])
		var lowest uint64 = math.MaxUint64
		var locs []uint64
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			spinnerNextFrame()
			needleOuter := j
			rangeSize := seeds[i+1]

			currBatchLen := batchSize
			if currBatchLen > int(rangeSize) {
				currBatchLen = int(rangeSize)
			}

			wg.Add(currBatchLen)
			for offset := 0; offset < currBatchLen; offset++ {
				go func(seed uint64) {
					needle := seed
					for k := 0; k < len(mapping); k++ {
						resp := walkMap(mapping, k, needle)
						needle = resp[0]
					}

					locs = append(locs, needle)
					wg.Done()
				}(needleOuter)

				needleOuter++
			}

			wg.Wait()

			lowestBatchLoc := slices.Min[[]uint64](locs)

			if lowestBatchLoc > 0 && lowestBatchLoc < lowest {
				lowest = lowestBatchLoc
			}
			locs = locs[:0]
			j += uint64(currBatchLen)
		}

		seedLowest = append(seedLowest, lowest)
	}

	return slices.Min(seedLowest)
}

type Mapping struct {
	Start       map[uint64]uint64
	End         map[uint64]uint64
	SourceStart uint64
	SourceEnd   uint64
	Length      uint64
}

func walkMap(builtMap [][]Mapping, idx int, needle uint64) []uint64 {
	var response []uint64

	mapping := builtMap[idx]
	for i := 0; i < len(mapping); i++ {
		rangeMap := mapping[i]

		if needle >= rangeMap.SourceStart && needle <= rangeMap.SourceEnd {
			offset := needle - rangeMap.SourceStart
			return []uint64{rangeMap.Start[rangeMap.SourceStart] + offset}
		}
	}

	return append(response, needle)
}

func buildMap(content string) ([]uint64, [][]Mapping) {
	blocks := strings.Split(content, "\n\n")

	output := make([][]Mapping, len(blocks))
	for i := 0; i < len(blocks); i++ {
		blockLines := strings.Split(blocks[i], "\n")

		blockMapping := make([]Mapping, 0)
		for j := 1; j < len(blockLines); j++ {
			lineParts := strings.Split(blockLines[j], " ")

			src := MustUInt64(lineParts[1])
			dst := MustUInt64(lineParts[0])
			rangeLength := MustUInt64(lineParts[2])

			blockMapping = append(blockMapping, Mapping{
				Start:       map[uint64]uint64{src: dst},
				End:         map[uint64]uint64{src + rangeLength - 1: dst + rangeLength - 1},
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

func processSeedBlock(block string) []uint64 {
	cleanBlock := strings.TrimPrefix(block, "seeds:")

	seeds := []uint64{}
	for _, seed := range strings.Split(cleanBlock, " ") {
		if num, err := strconv.ParseUint(seed, 10, 64); err == nil {
			seeds = append(seeds, num)
		}
	}

	return seeds
}

// ================================================
// Helper fuctions
// ================================================
func MustUInt64(target string) uint64 {
	if num, err := strconv.ParseUint(target, 10, 64); err == nil {
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
