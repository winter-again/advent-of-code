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

	scanner := bufio.NewScanner(file)
	// todo: maybe this should be []string instead and then do conversion to int
	// only when needed for calcs?
	var stones []int
	for scanner.Scan() {
		stoneStr := strings.Fields(scanner.Text())
		for _, stone := range stoneStr {
			stoneInt, _ := strconv.Atoi(stone)
			stones = append(stones, stoneInt)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	n := 25
	for i := 0; i < n; i++ {
		var newStones []int
		// todo: do copy instead?
		// todo: can't modify stones while iterating?
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(strconv.Itoa(stone))%2 == 0 {
				left, right := splitStone(stone)
				newStones = append(newStones, left, right)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}
	return len(stones)
}

func splitStone(stone int) (int, int) {
	stoneStr := strconv.Itoa(stone)
	left, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
	right, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
	return left, right
}

type stone struct {
	num  int
	next *stone
}

type stoneList struct {
	head *stone
}

func (sl *stoneList) addStone(stoneNum int) {
	st := &stone{
		num: stoneNum,
	}
	if sl.head == nil {
		sl.head = st
	} else {
		cur := sl.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = st
	}
}

func solvePart1LinkedList(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var st []int
	for scanner.Scan() {
		stoneStr := strings.Fields(scanner.Text())
		for _, stone := range stoneStr {
			stoneInt, _ := strconv.Atoi(stone)
			st = append(st, stoneInt)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	stones := stoneList{}
	for _, stone := range st {
		stones.addStone(stone)
	}

	stones.addStone(-1)

	num := len(st)
	n := 25
	for i := 0; i < n; i++ {
		cur := stones.head
		for cur.next != nil {
			if cur.num == -1 {
				break
			}

			if cur.num == 0 {
				cur.num = 1
				cur = cur.next
			} else if len(strconv.Itoa(cur.num))%2 == 0 {
				left, right := splitStone(cur.num)
				rStone := &stone{num: right}

				next := cur.next
				cur.num = left
				cur.next = rStone
				rStone.next = next

				cur = cur.next.next
				num++
			} else {
				cur.num = cur.num * 2024
				cur = cur.next
			}
		}
	}
	return num
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stones := make(map[int]int)
	for scanner.Scan() {
		stoneStr := strings.Fields(scanner.Text())
		for _, stone := range stoneStr {
			stoneInt, _ := strconv.Atoi(stone)
			stones[stoneInt]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	n := 75
	for i := 0; i < n; i++ {
		newStones := make(map[int]int, 100)
		for stone, ct := range stones {
			if stone == 0 {
				newStones[1] += ct
			} else if len(strconv.Itoa(stone))%2 == 0 {
				left, right := splitStone(stone)
				newStones[left] += ct
				newStones[right] += ct
			} else {
				newStones[stone*2024] += ct
			}
		}

		stones = newStones
	}

	sum := 0
	for _, ct := range stones {
		sum += ct
	}
	return sum
}
