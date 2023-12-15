package lib

import "testing"

const TestData = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestPart01(t *testing.T) {
	expect := 6440
	actual := Part01(TestData)

	if expect != actual {
		t.Fatalf("want %v, got %v", expect, actual)
	}
}

func TestPart02(t *testing.T) {
	expect := 5905
	actual := Part02(TestData)

	if expect != actual {
		t.Fatalf("want %v, got %v", expect, actual)
	}
}
