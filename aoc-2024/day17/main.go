package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	part := flag.Int("part", 1, "problem part")
	flag.Parse()
	fmt.Println("Solving part", *part)

	if *part == 1 {
		ans := solvePart1("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart1("./input.txt")
		fmt.Println("Answer:", ans)
	} else {
		// ans := solvePart2("./input_smpl_2.txt")
		// fmt.Println("Answer (sample):", ans)

		ans := solvePart2("./input.txt")
		fmt.Println("Answer:", ans)
	}
}

func solvePart1(input string) string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reg, prog := parseProg(file)
	res := compute(reg, prog)
	ans := make([]string, len(res))
	for i, a := range res {
		ans[i] = strconv.Itoa(a)
	}
	return strings.Join(ans, ",")
}

func compute(reg register, prog []int) []int {
	var out []int
	i := 0
	for i < len(prog) {
		opcode := prog[i]
		operand := prog[i+1]
		if opcode == 0 {
			c, err := combo(operand, &reg)
			if err != nil {
				log.Fatal(err)
			}
			reg.a = reg.a >> c
		} else if opcode == 1 {
			reg.b = reg.b ^ operand
		} else if opcode == 2 {
			c, err := combo(operand, &reg)
			if err != nil {
				log.Fatal(err)
			}
			reg.b = c % 8
		} else if opcode == 3 {
			if reg.a != 0 {
				i = operand
				continue
			}
		} else if opcode == 4 {
			reg.b = reg.b ^ reg.c
		} else if opcode == 5 {
			c, err := combo(operand, &reg)
			if err != nil {
				log.Fatal(err)
			}
			out = append(out, c%8)
		} else if opcode == 6 {
			c, err := combo(operand, &reg)
			if err != nil {
				log.Fatal(err)
			}
			reg.b = reg.a >> c
		} else if opcode == 7 {
			c, err := combo(operand, &reg)
			if err != nil {
				log.Fatal(err)
			}
			reg.c = reg.a >> c
		}
		i += 2
	}
	return out
}

func combo(op int, reg *register) (int, error) {
	if op >= 0 && op <= 3 {
		return op, nil
	} else if op == 4 {
		return reg.a, nil
	} else if op == 5 {
		return reg.b, nil
	} else if op == 6 {
		return reg.c, nil
	}
	return 0, fmt.Errorf("7 is reserved")
}

type register struct {
	a int
	b int
	c int
}

func parseProg(file *os.File) (register, []int) {
	scanner := bufio.NewScanner(file)
	scanner.Split(splitProg)

	scanner.Scan()
	regFields := strings.Split(scanner.Text(), "\n")
	regA, _ := strconv.Atoi(strings.TrimLeftFunc(regFields[0], func(r rune) bool {
		return !unicode.IsNumber(r)
	}))
	regB, _ := strconv.Atoi(strings.TrimLeftFunc(regFields[1], func(r rune) bool {
		return !unicode.IsNumber(r)
	}))
	regC, _ := strconv.Atoi(strings.TrimLeftFunc(regFields[2], func(r rune) bool {
		return !unicode.IsNumber(r)
	}))
	registers := register{
		a: regA,
		b: regB,
		c: regC,
	}

	scanner.Scan()
	progPre := strings.Split(strings.TrimPrefix(strings.TrimSpace(scanner.Text()), "Program: "), ",")
	prog := make([]int, len(progPre))
	for i, p := range progPre {
		out, _ := strconv.Atoi(p)
		prog[i] = out
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return registers, prog
}

func splitProg(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}
	return
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, prog := parseProg(file)

	// NOTE: transl
	// 2,4: regB = regA % 8 -> mod makes this have only 8 possibs ([0, 7]); keep only last 3 bits
	// 1,3: regB = regB ^ 3 -> flips the first 2 bits (bin(3) = 11)
	// ----------------
	// 7,5: regC = regA >> regB -> truncates first regB bits
	// ----------------
	// 0,3: regA = regA >> 3 -> only mod of A; truncates first 3 bits; only used for loop term after
	// ----------------
	// 4,1: regB = regB ^ regC -> flips some bits
	// 1,5: regB = regB ^ 5 -> flips first and third bits (bin(5) = 101)
	// ----------------
	// 5,5: out(regB % 8) -> print regB % 8
	// 3,0: if regA != 0 { jump to 0 } -> i.e., loop back to first instr (i=0)
	// eventually, regA will be 0 to end loop

	// NOTE: work backwards to find regA given constraints
	// term) init regA must be <= 7 in last iter for final regA = init regA >> 3 to set regA to 0 and term
	// note that bin(7) = 111 so 7 >> 3 = 0
	// out 0) to output the 0, want final regB % 8 = 0
	// thus *init* regA must allow regB % 8 to become 0 in this iter
	// try regA = 0...7 to see impact on regB and whether it satisfies regB % 8 = 0
	// from sampling I see *only init regA = 6 prints 0* (and 6 >> 3 = 0 for term)
	// out 3) thus, this loop ends with regA = 6, meaning init regA >> 3 = 6
	// so init regA = 6 << 3 = 48 -> however setting regA to 48 isn't enough
	// because it prints 0, 0 (got the ending 0 but not the 3)
	// so init regA can be 48...55 (8 possibs) -> these all give 6 when shifted with >> 3
	// aka 48 + 0...7
	// so have to determine which of those would work to print the proper "3,0"
	// 49 satisfies both

	try := []int{0}
	for l := 0; l < len(prog); l++ {
		var next []int
		for _, v := range try {
			for b := 0; b < 8; b++ {
				targA := v<<3 | b
				reg := register{a: targA}
				// log.Println("TRY:", try)
				// log.Println("COMPARE", compute(reg, prog), prog, prog[len(prog)-1-l:])
				if slices.Equal(compute(reg, prog), prog[len(prog)-1-l:]) {
					next = append(next, targA)
				}
			}
		}
		try = next
	}
	return slices.Min(try)
}
