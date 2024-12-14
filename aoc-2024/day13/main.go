package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
		ans := solvePart2("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart2("./input.txt")
		fmt.Println("Answer:", ans)
	}
}

func solvePart1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	machines := parseMachines(file, 1)

	tokens := 0
	for _, m := range machines {
		// note: assumes soln exists
		pressB := float64(m.a.y*m.prize.x-m.a.x*m.prize.y) / float64(m.a.y*m.b.x-m.a.x*m.b.y)
		pressA := (float64(m.prize.x) - float64(m.b.x)*pressB) / float64(m.a.x)
		if pressA == float64(int64(pressA)) && pressB == float64(int64(pressB)) {
			tokens += int(pressA*3) + int(pressB)
		}
	}
	return tokens
}

type coord struct {
	x int
	y int
}

type machine struct {
	a     coord
	b     coord
	prize coord
}

func parseMachines(file *os.File, part int) []machine {
	scanner := bufio.NewScanner(file)
	scanner.Split(splitMachines)

	var machines []machine
	for scanner.Scan() {
		var m machine
		fields := strings.Split(strings.TrimSpace(scanner.Text()), "\n")

		a := strings.Split(strings.TrimPrefix(fields[0], "Button A: "), ", ")
		_, aX, _ := strings.Cut(a[0], "+")
		aXInt, _ := strconv.Atoi(aX)
		_, aY, _ := strings.Cut(a[1], "+")
		aYInt, _ := strconv.Atoi(aY)
		m.a = coord{aXInt, aYInt}

		b := strings.Split(strings.TrimPrefix(fields[1], "Button B: "), ", ")
		_, bX, _ := strings.Cut(b[0], "+")
		bXInt, _ := strconv.Atoi(bX)
		_, bY, _ := strings.Cut(b[1], "+")
		bYInt, _ := strconv.Atoi(bY)
		m.b = coord{bXInt, bYInt}

		p := strings.Split(strings.TrimPrefix(fields[2], "Prize: "), ", ")
		_, pX, _ := strings.Cut(p[0], "=")
		pXInt, _ := strconv.Atoi(pX)
		_, pY, _ := strings.Cut(p[1], "=")
		pYInt, _ := strconv.Atoi(pY)

		if part == 2 {
			pXInt = pXInt + 10000000000000
			pYInt = pYInt + 10000000000000
		}

		m.prize = coord{pXInt, pYInt}

		machines = append(machines, m)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return machines
}

func splitMachines(data []byte, atEOF bool) (advance int, token []byte, err error) {
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

	machines := parseMachines(file, 2)

	tokens := 0
	for _, m := range machines {
		// note: assumes soln exists
		pressB := float64(m.a.y*m.prize.x-m.a.x*m.prize.y) / float64(m.a.y*m.b.x-m.a.x*m.b.y)
		pressA := (float64(m.prize.x) - float64(m.b.x)*pressB) / float64(m.a.x)
		if pressA == float64(int64(pressA)) && pressB == float64(int64(pressB)) {
			tokens += int(pressA*3) + int(pressB)
		}
	}
	return tokens
}
