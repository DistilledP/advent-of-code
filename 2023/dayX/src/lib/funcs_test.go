package lib

import "testing"

const TestData = ``

func TestPart01(t *testing.T) {
	t.Skip("Part01 is skipped")
	expect := 33
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
