package lib

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

/**
Benchmark results:
==================
./src/lib: go test -benchtime=5s -cover -benchmem -bench=.

BenchmarkPart01-10               4348591              1371 ns/op            1512 B/op         35 allocs/op
BenchmarkPart02-10                132846             44374 ns/op           66066 B/op        677 allocs/op
BenchmarkPart02_v2-10            2038118              2949 ns/op             126 B/op          8 allocs/op
BenchmarkPart02_v3-10             573256             10401 ns/op            1918 B/op         36 allocs/op
PASS
coverage: 93.9% of statements
*/

// Part01 Combined value is: 53080
func Part01(input string) int {
	split := strings.Split(input, "\n")
	total := 0

	for i := 0; i < len(split); i++ {
		first := 0
		second := first

		for i2 := 0; i2 < len(split[i]); i2++ {
			if num, ok := strconv.Atoi(string(split[i][i2])); ok == nil {
				if first == 0 {
					first = num
				}
				second = num
			}
		}

		if combined, ok := strconv.Atoi(fmt.Sprintf("%d%d", first, second)); ok == nil {
			total += combined
		}
	}

	return total
}

// Part02 Combined value is: 53268
func Part02(input string) int {
	split := strings.Split(input, "\n")

	total := 0
	for i := 0; i < len(split); i++ {
		expr1, err := regexp.Compile(`[0-9]|(one|two|three|four|five|six|seven|eight|nine)`)
		if err != nil {
			log.Fatalf("failed to compile regex: %s", err)
		}

		firstMatch := expr1.FindString(split[i])

		var reversedStr string
		for i2 := len(split[i]) - 1; i2 >= 0; i2-- {
			reversedStr += string(split[i][i2])
		}

		var foo, lastMatch string
		for i3 := 0; i3 < len(reversedStr); i3++ {
			foo = string(reversedStr[i3]) + foo
			lastMatch = expr1.FindString(foo)
			if len(lastMatch) > 0 {
				break
			}
		}

		first := convNumStr(firstMatch)
		last := convNumStr(lastMatch)

		combined := fmt.Sprintf("%d%d", first, last)

		if num, err := strconv.Atoi(combined); err == nil {
			total += num
		}
	}

	return total
}

// Combined value is: 53268
func Part02_v2(input string) int {
	lines := strings.Split(input, "\n")
	total := 0

	for i := 0; i < len(lines); i++ {
		total += totalLine(lines[i])
	}

	return total
}

func totalLine(line string) int {
	var first, last int
	for i := 0; i < len(line); i++ {
		if first = numberInLine(line[i:], strings.HasPrefix); first > 0 {
			break
		}
	}

	for i := len(line); i >= 0; i-- {
		if last = numberInLine(line[:i], strings.HasSuffix); last > 0 {
			break
		}
	}

	if num, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last)); err == nil {
		return num
	}

	return 0
}

// Combined value is: 53268
func Part02_v3(input string) int {
	lines := strings.Split(input, "\n")
	total := 0

	for i := 0; i < len(lines); i++ {
		total += totalLine_v2(lines[i])
	}

	return total
}

func totalLine_v2(line string) int {
	first := <-getFirst(line)
	last := <-getLast(line)

	if num, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last)); err == nil {
		return num
	}

	return 0
}

func getFirst(line string) <-chan int {
	ch := make(chan int)

	run := func() {
		defer close(ch)
		for i := 0; i < len(line); i++ {
			if first := numberInLine(line[i:], strings.HasPrefix); first > 0 {
				ch <- first
				break
			}
		}
	}

	go run()
	return ch
}

func getLast(line string) <-chan int {
	ch := make(chan int)

	run := func() {
		defer close(ch)
		for i := len(line); i >= 0; i-- {
			if last := numberInLine(line[:i], strings.HasSuffix); last > 0 {
				ch <- last
				break
			}
		}
	}

	go run()

	return ch
}

type cmpFunc = func(string, string) bool

func numberInLine(line string, cmp cmpFunc) int {
	numbers := []string{
		"0", "zero",
		"1", "one",
		"2", "two",
		"3", "three",
		"4", "four",
		"5", "five",
		"6", "six",
		"7", "seven",
		"8", "eight",
		"9", "nine",
	}

	var numVal int
	for i := 0; i < len(numbers); i++ {
		if cmp(line, numbers[i]) {
			numVal = convNumStr(numbers[i])
			break
		}
	}

	return numVal
}

func convNumStr(v string) int {
	switch v {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		if num, err := strconv.Atoi(v); err == nil {
			return num
		}
	}

	return 0
}
