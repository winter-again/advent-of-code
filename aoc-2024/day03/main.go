package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

		ans = solvePart1NoRegex("./input_smpl.txt")
		fmt.Println("Answer (sample) w/ no regex:", ans)

		ans = solvePart1("./input.txt")
		fmt.Println("Answer:", ans)

		ans = solvePart1NoRegex("./input.txt")
		fmt.Println("Answer w/ no regex:", ans)

		ans = solvePart1NoRegex("./input.txt")
		fmt.Println("Answer w/ no regex:", ans)
	} else {
		ans := solvePart2("./input_smpl_2.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart2("./input.txt")
		fmt.Println("Answer:", ans)

		ans = solvePart2NoRegex("./input_smpl_2.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart2NoRegex("./input.txt")
		fmt.Println("Answer:", ans)
	}
}

func solvePart1(input string) int {
	mem, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	memStr := string(mem)
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	parsedMem := re.FindAllString(memStr, -1)

	res := make([]int, len(parsedMem))
	for i, op := range parsedMem {
		res[i] = mul(op)
	}
	return sum(res)
}

func mul(expr string) int {
	re := regexp.MustCompile(`\d+,\d+`)
	pair := re.FindString(expr)
	nums := strings.Split(pair, ",")

	x, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatal(err)
	}
	return x * y
}

func sum(s []int) int {
	sum := 0
	for _, x := range s {
		sum += x
	}
	return sum
}

func solvePart1NoRegex(input string) int {
	mem, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	memStr := string(mem)
	res := 0
	for i := 0; i < len(memStr)-8; i++ {
		if memStr[i] == 'm' {
			if memStr[i+1] == 'u' && memStr[i+2] == 'l' && memStr[i+3] == '(' {
				i += 4
				var x int
				var y int
				if memStr[i] >= 48 && memStr[i] <= 57 {
					x = parseNumber(memStr, &i)
					if memStr[i] == ',' {
						i += 1
						y = parseNumber(memStr, &i)
						if memStr[i] == ')' {
							res += x * y
						}
					}
				}
			}
		}
	}
	return res
}

func parseNumber(memStr string, i *int) int {
	var sb strings.Builder
	for {
		if !(memStr[*i] >= 48 && memStr[*i] <= 57) {
			break
		}
		sb.WriteString(string(memStr[*i]))
		*i++
	}

	num, err := strconv.Atoi(sb.String())
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func solvePart2(input string) int {
	mem, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	memStr := string(mem)
	re := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
	parsedMem := re.FindAllString(memStr, -1)

	var dont bool
	doSwitch := "do()"
	dontSwitch := "don't()"

	var validOps []string
	for _, op := range parsedMem {
		if op == dontSwitch {
			dont = true
		} else if op == doSwitch {
			dont = false
		} else {
			if !dont {
				validOps = append(validOps, op)
			}
		}
	}

	res := make([]int, len(validOps))
	for i, op := range validOps {
		res[i] = mul(op)
	}
	return sum(res)
}

func solvePart2NoRegex(input string) int {
	mem, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	memStr := string(mem)
	res := 0
	ok := true
	for i := 0; i < len(memStr)-8; i++ {
		if memStr[i] == 'd' && memStr[i+1] == 'o' && memStr[i+2] == 'n' && memStr[i+3] == '\'' && memStr[i+4] == 't' {
			ok = false
			i += 7
		}

		if memStr[i] == 'd' && memStr[i+1] == 'o' && memStr[i+2] == '(' && memStr[i+3] == ')' {
			ok = true
			i += 4
		}

		if ok {
			if memStr[i] == 'm' {
				if memStr[i+1] == 'u' && memStr[i+2] == 'l' && memStr[i+3] == '(' {
					i += 4
					var x int
					var y int
					if memStr[i] >= 48 && memStr[i] <= 57 {
						x = parseNumber(memStr, &i)
						if memStr[i] == ',' {
							i += 1
							y = parseNumber(memStr, &i)
							if memStr[i] == ')' {
								res += x * y
							}
						}
					}
				}
			}
		}
	}
	return res
}
