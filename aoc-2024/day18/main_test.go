package main

import (
	"testing"

	"github.com/winter-again/advent-of-code/aoc-2024/util"
)

func TestPart1(t *testing.T) {
	tests := []util.TestCase{
		{
			Name:   "Part 1 (sample)",
			Input:  "./input_smpl.txt",
			Expect: 22,
		},
		{
			Name:   "Part 1",
			Input:  "./input.txt",
			Expect: 348,
		},
	}

	t.Run(tests[0].Name, func(t *testing.T) {
		if got := solvePart1(tests[0].Input, true); got != tests[0].Expect {
			t.Errorf("%s: expected %d but got %d", tests[0].Name, tests[0].Expect, got)
		}
	})
	t.Run(tests[1].Name, func(t *testing.T) {
		if got := solvePart1(tests[1].Input, false); got != tests[1].Expect {
			t.Errorf("%s: expected %d but got %d", tests[1].Name, tests[1].Expect, got)
		}
	})
}

func TestPart2(t *testing.T) {
	tests := []util.TestCaseDay18Part2{
		{
			Name:   "Part 2 (sample)",
			Input:  "./input_smpl.txt",
			Sample: true,
			Expect: "6,1",
		},
		{
			Name:   "Part 2",
			Input:  "./input.txt",
			Sample: false,
			Expect: "54,44",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if got := solvePart2(test.Input, test.Sample); got != test.Expect {
				t.Errorf("%s: expected %s but got %s", test.Name, test.Expect, got)
			}
		})
	}
}
