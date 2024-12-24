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
			Expect: 7,
		},
		{
			Name:   "Part 1",
			Input:  "./input.txt",
			Expect: 1062,
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
	test1 := util.TestCaseDay17Part1{
		Name:   "Part 2 (sample)",
		Input:  "./input_smpl.txt",
		Expect: "co,de,ka,ta",
	}
	test2 := util.TestCaseDay17Part1{
		Name:   "Part 2",
		Input:  "./input.txt",
		Expect: "bz,cs,fx,ms,oz,po,sy,uh,uv,vw,xu,zj,zm",
	}

	t.Run(test1.Name, func(t *testing.T) {
		if got := solvePart2(test1.Input); got != test1.Expect {
			t.Errorf("%s: expected %s but got %s", test1.Name, test1.Expect, got)
		}
	})

	t.Run(test2.Name, func(t *testing.T) {
		if got := solvePart2(test2.Input); got != test2.Expect {
			t.Errorf("%s: expected %s but got %s", test2.Name, test2.Expect, got)
		}
	})
}
