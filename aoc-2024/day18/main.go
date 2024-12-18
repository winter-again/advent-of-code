package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	part := flag.Int("part", 1, "problem part")
	flag.Parse()
	fmt.Println("Solving part", *part)

	if *part == 1 {
		ans := solvePart1("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		// ans = solvePart1("./input.txt")
		// fmt.Println("Answer:", ans)
	} else {
		ans := solvePart2("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		// ans = solvePart2("./input.txt")
		// fmt.Println("Answer:", ans)
	}
}

func solvePart1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return 0
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return 0
}
