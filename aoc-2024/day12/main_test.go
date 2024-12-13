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
			Expect: 140,
		},
		{
			Name:   "Part 1 (sample 2)",
			Input:  "./input_smpl_2.txt",
			Expect: 1930,
		},
		{
			Name:   "Part 1",
			Input:  "./input.txt",
			Expect: 1344578,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if got := solvePart1(test.Input); got != test.Expect {
				t.Errorf("%s: expected %d but got %d", test.Name, test.Expect, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []util.TestCase{
		{
			Name:   "Part 2 (sample)",
			Input:  "./input_smpl.txt",
			Expect: 80,
		},
		{
			Name:   "Part 1 (sample 2)",
			Input:  "./input_smpl_2.txt",
			Expect: 1206,
		},
		{
			Name:   "Part 2",
			Input:  "./input.txt",
			Expect: 814302,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if got := solvePart2(test.Input); got != test.Expect {
				t.Errorf("%s: expected %d but got %d", test.Name, test.Expect, got)
			}
		})
	}
}
