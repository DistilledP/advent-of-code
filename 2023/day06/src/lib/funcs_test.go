package lib

import (
	"testing"
)

const TestData = `Time:      7  15   30
Distance:  9  40  200`

func TestPart01(t *testing.T) {
	expect := 288
	actual := Part01(TestData)

	if expect != actual {
		t.Fatalf("want %v, got %v", expect, actual)
	}
}

func TestPart02(t *testing.T) {
	t.Skip("Part02 is skipped")
	expect := 33
	actual := Part02(TestData)

	if expect != actual {
		t.Fatalf("want %v, got %v", expect, actual)
	}
}

type testCase struct {
	held       int
	raceLength int
	expected   int
}

func TestCalculateDistanceTravelled(t *testing.T) {
	testCases := []testCase{
		{
			held:       1,
			raceLength: 7,
			expected:   6,
		},
		{
			held:       2,
			raceLength: 7,
			expected:   10,
		},
		{
			held:       3,
			raceLength: 7,
			expected:   12,
		},
		{
			held:       4,
			raceLength: 7,
			expected:   12,
		},
		{
			held:       5,
			raceLength: 7,
			expected:   10,
		},
		{
			held:       6,
			raceLength: 7,
			expected:   6,
		},
		{
			held:       7,
			raceLength: 7,
			expected:   0,
		},
	}

	for tcId, tc := range testCases {
		actual := calculateDistanceTravelled(tc.held, tc.raceLength)
		if actual != tc.expected {
			t.Fatalf("[%d] want %v, got %v", tcId+1, tc.expected, actual)
		}
	}
}
