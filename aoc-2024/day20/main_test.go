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
			Expect: 14 + 14 + 2 + 4 + 2 + 3 + 1 + 1 + 1 + 1 + 1,
		},
		{
			Name:   "Part 1",
			Input:  "./input.txt",
			Expect: 1355,
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
	tests := []util.TestCase{
		{
			Name:   "Part 2 (sample)",
			Input:  "./input_smpl.txt",
			Expect: 32 + 31 + 29 + 39 + 25 + 23 + 20 + 19 + 12 + 14 + 12 + 22 + 4 + 3,
		},
		{
			Name:   "Part 2",
			Input:  "./input.txt",
			Expect: 1007335,
		},
	}

	t.Run(tests[0].Name, func(t *testing.T) {
		if got := solvePart2(tests[0].Input, true); got != tests[0].Expect {
			t.Errorf("%s: expected %d but got %d", tests[0].Name, tests[0].Expect, got)
		}
	})
	t.Run(tests[1].Name, func(t *testing.T) {
		if got := solvePart2(tests[1].Input, false); got != tests[1].Expect {
			t.Errorf("%s: expected %d but got %d", tests[1].Name, tests[1].Expect, got)
		}
	})
}
