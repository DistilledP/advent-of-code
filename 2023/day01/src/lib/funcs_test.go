package lib

import (
	"testing"
)

var result int

func Test_part01_works(t *testing.T) {
	const testInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	expect := 142
	actual := Part01(testInput)

	if expect != actual {
		t.Fatalf("want %v, got %v", expect, actual)
	}
}

func Test_part02_works(t *testing.T) {
	const testInput = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	expect := 281
	actual := Part02(testInput)

	if expect != actual {
		t.Fatalf("want %v, got %v", expect, actual)
	}
}

func Test_part02v2_works(t *testing.T) {
	const testInput = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	expect := 281
	actual := Part02_v2(testInput)

	if expect != actual {
		t.Fatalf("want %v, got %v", expect, actual)
	}
}

func BenchmarkPart01(b *testing.B) {
	const testInput = `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	var r int
	for n := 0; n < b.N; n++ {
		r = Part01(testInput)
	}

	result = r
}

func BenchmarkPart02(b *testing.B) {
	const testInput = `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	var r int
	for n := 0; n < b.N; n++ {
		r = Part02(testInput)
	}

	result = r
}

func BenchmarkPart02_v2(b *testing.B) {
	const testInput = `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	var r int
	for n := 0; n < b.N; n++ {
		r = Part02_v2(testInput)
	}

	result = r
}

func BenchmarkPart02_v3(b *testing.B) {
	const testInput = `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	var r int
	for n := 0; n < b.N; n++ {
		r = Part02_v3(testInput)
	}

	result = r
}
