package lib

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Part01: answer 535235
func Part01(content string) int {
	regExNumbers, err := regexp.Compile(`[\d]+`)
	if err != nil {
		log.Fatalf("failed to compile regex: %s\n", err)
	}

	regExSymbols, err := regexp.Compile(`[#$%&*+\-/=@]`)
	if err != nil {
		log.Fatalf("failed to compile symbols regex: %s\n", err)
	}

	var rollingVal int

	lines := strings.Split(content, "\n")
	for i := 0; i < len(lines); i++ {
		numbers := regExNumbers.FindAllStringSubmatchIndex(lines[i], -1)
		if len(numbers) > 0 {
			for i2 := 0; i2 < len(numbers); i2++ {
				var start, end int

				if numbers[i2][0] > 1 {
					start = numbers[i2][0] - 1
				}

				if numbers[i2][1] <= len(lines[i]) {
					end = numbers[i2][1] + 1
					if end > len(lines[i]) {
						end = len(lines[i]) - 1
					}
				}

				above := ""
				if (i - 1) >= 0 {
					above = lines[i-1][start:end]
				}

				inline := lines[i][start:end]

				below := ""
				if (i + 1) < len(lines) {
					below = lines[i+1][start:end]
				}

				hasMatch := regExSymbols.MatchString(above) || regExSymbols.MatchString(inline) || regExSymbols.MatchString(below)

				debugLog(lines[i], numbers[i2][0], numbers[i2][1], lines[i][numbers[i2][0]:numbers[i2][1]], "above", above, "inline", inline, "below", below, "hasMatch", hasMatch)

				if hasMatch {
					if num, err := strconv.Atoi(lines[i][numbers[i2][0]:numbers[i2][1]]); err == nil {
						rollingVal += num
					}
				}
			}
		}
	}

	return rollingVal
}

const (
	ASCII_FULLSTOP = 55
	ASCII_ASTERIX  = 42
)

type loc struct {
	X, Y int
}

type gear struct {
	Loc                loc
	UpperVal, LowerVal int
}

type parsedContent struct {
	Grid  [][]byte
	Gears []gear
}

/*
	 Part02:
		   58801205 is too low, suspect we're missing inline
		   72030072 is also too low, though this doesn't include complete inline (manually calculated)
		   74166837 is also not right, manually added the fill inline X*X to the previous
*/
func Part02(content string) int {
	grid := parseContent(content)

	// for i := 0; i < len(grid.Grid); i++ {
	// 	debugLog(string(grid.Grid[i]), len(grid.Grid[i]))
	// }

	runningTotal := 0
	for i := 0; i < len(grid.Gears); i++ {
		runningTotal += grid.Gears[i].UpperVal * grid.Gears[i].LowerVal
	}

	// gearOffset := len(grid.Gears) - 16

	// debugLog("upper: ", grid.Gears[gearOffset].UpperVal, "lower: ", grid.Gears[gearOffset].LowerVal)

	return runningTotal
}

func parseContent(content string) parsedContent {
	rows := strings.Split(content, "\n")
	output := parsedContent{
		Grid: make([][]byte, len(rows)+2),
	}

	// Populate with "empty" values
	emptyLine := bytes.Repeat([]byte("."), maxLen(rows)+2)
	for i := range output.Grid {
		output.Grid[i] = emptyLine
	}

	// Populate the grid with data
	for y := 0; y < len(rows); y++ {
		line := rows[y]
		yLoc := y + 1
		output.Grid[yLoc] = make([]byte, len(line)+2)

		for x := 0; x < len(line); x++ {
			xLoc := x + 1
			output.Grid[yLoc][xLoc] = line[x]
		}
	}

	// Parse Symbols
	for y := 1; y < len(output.Grid)-1; y++ {
		for x := 1; x < len(output.Grid)-1; x++ {
			switch output.Grid[y][x] {
			case ASCII_ASTERIX:
				upperY := y - 1
				lowerY := y + 1

				leftX := x - 1
				rightX := x + 2

				upper := output.Grid[upperY][leftX:rightX]
				fullUpper := output.Grid[upperY]

				lower := output.Grid[lowerY][leftX:rightX]
				fullLower := output.Grid[lowerY]

				inline := output.Grid[y][leftX:rightX]
				fullInline := output.Grid[y]

				upperHasNum := bytes.ContainsAny(upper, "0123456789")
				lowerHasNum := bytes.ContainsAny(lower, "0123456789")
				inlineHasNum := bytes.ContainsAny(inline, "0123456789")

				// if inlineHasNum {
				// debugLog(string(inline), string(fullInline))
				// }

				// suspect we're missing "inline"

				// logoutput := fmt.Sprintf(
				// 	`x: %v, upperY: %d, lowerY: %d, leftX: %d, rightX: %d, fullUpper: %s, upper: %s, upperHasNum: %v, fullLower: %s, lower: %s, lowerHasNum: %v`,
				// 	x, upperY, lowerY, leftX, rightX, string(fullUpper), string(upper), upperHasNum, string(fullLower), string(lower), lowerHasNum,
				// )

				// debugLog(logoutput)

				if upperHasNum && lowerHasNum || upperHasNum && inlineHasNum || inlineHasNum && lowerHasNum || lowerHasNum && !upperHasNum && !lowerHasNum {
					newGear := gear{
						Loc: loc{
							X: x,
							Y: y,
						},
					}

					// Now we need to get the numbers
					// Probably better to find the bounds, pass the line into a function and the starting location.
					// traverse left/right until a "." is found and splice them together.
					upperValue := getValue(fullUpper, x)
					lowerValue := getValue(fullLower, x)
					inlineValue := getValue(fullInline, x)

					if inlineHasNum {
						debugLog(strings.Trim(fmt.Sprintf("%s\t%s\t%s", strings.TrimSpace(inlineValue), strings.TrimSpace(upperValue), strings.TrimSpace(lowerValue)), ""))
					}

					if upperVal, err := strconv.Atoi(upperValue); err == nil {
						newGear.UpperVal = upperVal
					}

					if lowerVal, err := strconv.Atoi(lowerValue); err == nil {
						newGear.LowerVal = lowerVal
					}

					output.Gears = append(output.Gears, newGear)
				}
			}
		}
	}

	return output
}

func getValue(value []byte, x int) string {
	var leftSide, rightSide string

	pos := x
	for pos > 0 {
		if pos < x && string(value[pos]) == "." {
			break
		}

		if string(value[pos]) != "." && string(value[pos]) != "*" && value[pos] != 0 {
			leftSide = fmt.Sprintf("%s%s", string(value[pos]), leftSide)
		}
		pos--
	}

	posX := x + 1
	for posX < len(value) {
		if posX > x && string(value[posX]) == "." {
			break
		}
		if string(value[posX]) != "." && string(value[pos]) != "*" && value[posX] != 0 {
			rightSide = fmt.Sprintf("%s%s", rightSide, string(value[posX]))
		}

		posX++
	}

	return fmt.Sprintf("%s%s", leftSide, rightSide)
}

func maxLen(items []string) int {
	maxLen := 0

	for i := 0; i < len(items); i++ {
		if len(items[i]) > maxLen {
			maxLen = len(items[i])
		}
	}

	return maxLen
}

func debugLog(items ...any) {
	if os.Getenv("DEBUG") == "1" {
		fmt.Println(items...)
	}
}
