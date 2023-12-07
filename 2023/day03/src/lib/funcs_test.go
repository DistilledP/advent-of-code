package lib

import (
	"log"
	"os"
	"testing"
)

const TestData = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPart01(t *testing.T) {
	t.Skip()
	expect := 4361
	actual := Part01(string(TestData))

	if expect != actual {
		t.Fatalf("want %v, got %v", expect, actual)
	}

}

func TestPart02(t *testing.T) {

	realData, err := os.ReadFile("./../../input/input")
	if err != nil {
		log.Fatalf("failed to read input file: %s", err)
	}

	expect := 467835
	// actual := Part02(TestData)
	actual := Part02(string(realData))

	if expect != actual {
		t.Fatalf("want %v, got %v", expect, actual)
	}
}
