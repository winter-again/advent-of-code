package main

import (
	"testing"

	"github.com/winter-again/advent-of-code/aoc-2024/util"
)

func TestPart1(t *testing.T) {
	tests := []util.TestCaseDay17Part1{
		{
			Name:   "Part 1 (sample)",
			Input:  "./input_smpl.txt",
			Expect: "4,6,3,5,6,3,5,2,1,0",
		},
		{
			Name:   "Part 1",
			Input:  "./input.txt",
			Expect: "1,5,0,5,2,0,1,3,5",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if got := solvePart1(test.Input); got != test.Expect {
				t.Errorf("%s: expected %s but got %s", test.Name, test.Expect, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []util.TestCase{
		{
			Name:   "Part 2 (sample)",
			Input:  "./input_smpl_2.txt",
			Expect: 117440,
		},
		{
			Name:   "Part 2",
			Input:  "./input.txt",
			Expect: 236581108670061,
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
