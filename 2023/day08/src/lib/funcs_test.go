package lib

import "testing"

const TestData = ``

type testCase struct {
	testData string
	expected int
}

func TestPart01(t *testing.T) {
	testCases := []testCase{
		{
			testData: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			expected: 2,
		},
		{
			testData: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
			expected: 6,
		},
	}

	for i, tc := range testCases {
		actual := Part01(tc.testData)

		if tc.expected != actual {
			t.Logf("[%d] want %v, got %v", i+1, tc.expected, actual)
			t.Fail()
		}
	}
}

func TestPart02(t *testing.T) {
	testCases := []testCase{
		{
			testData: `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`,
			expected: 6,
		},
	}

	for i, tc := range testCases {
		actual := Part02(tc.testData)

		if tc.expected != actual {
			t.Logf("[%d] want %v, got %v", i+1, tc.expected, actual)
			t.Fail()
		}
	}
}
