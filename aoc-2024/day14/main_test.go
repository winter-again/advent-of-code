package main

import (
	"testing"

	"github.com/winter-again/advent-of-code/aoc-2024/util"
)

func TestPart1(t *testing.T) {
	tests := []util.TestCaseDay14{
		{
			Name:   "Part 1 (sample)",
			Input:  "./input_smpl.txt",
			NRows:  7,
			NCols:  11,
			Expect: 12,
		},
		{
			Name:   "Part 1",
			Input:  "./input.txt",
			NRows:  103,
			NCols:  101,
			Expect: 218619324,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if got := solvePart1(test.Input, test.NRows, test.NCols); got != test.Expect {
				t.Errorf("%s: expected %d but got %d", test.Name, test.Expect, got)
			}
			if got := solvePart1Modulo(test.Input, test.NRows, test.NCols); got != test.Expect {
				t.Errorf("%s: expected %d but got %d", test.Name, test.Expect, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []util.TestCaseDay14{
		{
			Name:   "Part 2",
			Input:  "./input.txt",
			NRows:  103,
			NCols:  101,
			Expect: 6446,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if got := solvePart2(test.Input, test.NRows, test.NCols); got != test.Expect {
				t.Errorf("%s: expected %d but got %d", test.Name, test.Expect, got)
			}
		})
	}
}
