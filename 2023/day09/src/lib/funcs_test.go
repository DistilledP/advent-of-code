package lib

import "testing"

const TestData = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestPart01(t *testing.T) {
	var expect int64 = 114
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
