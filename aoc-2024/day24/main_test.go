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
			Expect: 4,
		},
		{
			Name:   "Part 1 (sample 2)",
			Input:  "./input_smpl_2.txt",
			Expect: 2024,
		},
		{
			Name:   "Part 1",
			Input:  "./input.txt",
			Expect: 45121475050728,
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
	tests := []util.TestCaseDay17Part1{
		// NOTE: retrospective, but I think didn't include this test case since approach requires
		// manual intervention/fixing as you progressively fix
		// {
		// 	Name:   "Part 2 (sample)",
		// 	Input:  "./input_smpl_3.txt",
		// 	Expect: "z00,z01,z02,z05",
		// },
		{
			Name:   "Part 2",
			Input:  "./input.txt",
			Expect: "gqp,hsw,jmh,mwk,qgd,z10,z18,z33",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if got := solvePart2(test.Input); got != test.Expect {
				t.Errorf("%s: expected %s but got %s", test.Name, test.Expect, got)
			}
		})
	}
}
