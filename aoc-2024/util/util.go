package util

type TestCase struct {
	Name   string
	Input  string
	Expect int
}

type TestCaseDay14 struct {
	Name   string
	Input  string
	NRows  int
	NCols  int
	Expect int
}

type TestCaseDay17Part1 struct {
	Name   string
	Input  string
	Expect string
}

type TestCaseDay18Part2 struct {
	Name   string
	Input  string
	Sample bool
	Expect string
}
