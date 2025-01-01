package main

import (
	"bufio"
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

		ans = solvePart1("./input.txt")
		fmt.Println("Answer:", ans)
	} else {
		ans := solvePart2("./input_smpl.txt")
		fmt.Println("Answer (sample):", ans)

		ans = solvePart2("./input.txt")
		fmt.Println("Answer:", ans)
	}
}

// NOTE: this is very slow
func solvePart1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	codes := parseCodes(file)
	sum := 0
	re := regexp.MustCompile("[0-9]+")
	for _, code := range codes {
		bot1Moves := expandMoves(code, numPad)

		var bot2Moves [][]string
		for _, seq := range bot1Moves {
			bot2Moves = append(bot2Moves, expandMoves(seq, dirPad)...)
		}

		var yourMoves [][]string
		for _, seq := range bot2Moves {
			yourMoves = append(yourMoves, expandMoves(seq, dirPad)...)
		}

		minLen := len(yourMoves[0])
		for _, seq := range yourMoves {
			if len(seq) < minLen {
				minLen = len(seq)
			}
		}

		codeNum, _ := strconv.Atoi(strings.Join(re.FindAllString(strings.Join(code, ""), -1), ""))
		cplx := minLen * codeNum
		sum += cplx
	}
	return sum
}

func expandMoves(code []string, keys map[string]pos) [][]string {
	var moves [][][]string
	var allMoves [][]string
	cur := keys["A"]
	for _, c := range code {
		var steps []string
		next := keys[c]
		diff := pos{next.i - cur.i, next.j - cur.j}

		if next == cur {
			steps = append(steps, "A")
		} else {
			var move string
			if diff.i < 0 {
				move = "^"
			} else if diff.i > 0 {
				move = "v"
			}
			for i := 0; i < absInt(diff.i); i++ {
				steps = append(steps, move)
			}

			if diff.j < 0 {
				move = "<"
			} else if diff.j > 0 {
				move = ">"
			}
			for j := 0; j < absInt(diff.j); j++ {
				steps = append(steps, move)
			}
		}

		stepPerms := permutations(steps)

		var validMoves [][]string
		for _, perm := range stepPerms {
			ok := true
			before := cur
			for _, p := range perm {
				d := disp[p]
				after := pos{before.i + d.i, before.j + d.j}
				if after == keys["X"] || !keyExists(after, keys) {
					ok = false
					break
				}
				before = after
			}

			if ok {
				if len(perm) == 1 && perm[0] == "A" {
					validMoves = append(validMoves, perm)
				} else {
					perm = append(perm, "A")
					validMoves = append(validMoves, perm)
				}
			}
		}
		moves = append(moves, validMoves)
		cur = next
	}

	allMoves = cartesian(moves...)
	return allMoves
}

func keyExists(loc pos, keys map[string]pos) bool {
	for _, val := range keys {
		if val == loc {
			return true
		}
	}
	return false
}

func cartesian(slices ...[][]string) [][]string {
	prod := slices[0]
	for _, s := range slices[1:] {
		var tmp [][]string
		for _, item1 := range prod {
			for _, item2 := range s {
				combined := append([]string{}, item1...)
				combined = append(combined, item2...)
				tmp = append(tmp, combined)
			}
		}
		prod = tmp
	}
	return prod
}

func permutations(s []string) [][]string {
	var perms [][]string
	var helper func([]string, int)
	seen := make(map[string]bool)

	helper = func(s []string, n int) {
		if n == 1 {
			tmp := make([]string, len(s))
			copy(tmp, s)
			key := strings.Join(tmp, "")
			if !seen[key] {
				perms = append(perms, tmp)
				seen[key] = true
			}
		} else {
			for i := 0; i < n; i++ {
				helper(s, n-1)
				if n%2 == 1 {
					tmp := s[i]
					s[i] = s[n-1]
					s[n-1] = tmp
				} else {
					tmp := s[0]
					s[0] = s[n-1]
					s[n-1] = tmp
				}
			}
		}
	}
	helper(s, len(s))
	return perms
}

type pos struct {
	i int
	j int
}

var numPad = map[string]pos{
	"X": {3, 0},
	"0": {3, 1},
	"A": {3, 2},
	"1": {2, 0},
	"2": {2, 1},
	"3": {2, 2},
	"4": {1, 0},
	"5": {1, 1},
	"6": {1, 2},
	"7": {0, 0},
	"8": {0, 1},
	"9": {0, 2},
}

var dirPad = map[string]pos{
	"<": {1, 0},
	"v": {1, 1},
	">": {1, 2},
	"X": {0, 0},
	"^": {0, 1},
	"A": {0, 2},
}

var disp = map[string]pos{
	"^": {-1, 0},
	"v": {1, 0},
	"<": {0, -1},
	">": {0, 1},
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseCodes(file *os.File) [][]string {
	nCodes := 5
	codes := make([][]string, nCodes)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		codes[i] = strings.Split(scanner.Text(), "")
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return codes
}

func solvePart2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	codes := parseCodes(file)
	sum := 0
	re := regexp.MustCompile("[0-9]+")
	cache := make(map[pairDepth]int)
	n := 25 // num of intermediate dirPad robots; n = 2 would give answer to part 1
	for _, code := range codes {
		bot1Moves := expandMoves(code, numPad)
		minLen := 9999999999999 // has to be big enough such that computed values guaranteed to be lower
		for _, seq := range bot1Moves {
			seq = append([]string{"A"}, seq...)
			l := 0
			for i := 0; i < len(seq)-1; i++ {
				a := seq[i]
				b := seq[i+1]
				l += dist(a, b, n, cache)
			}
			minLen = min(minLen, l)
		}
		codeNum, _ := strconv.Atoi(strings.Join(re.FindAllString(strings.Join(code, ""), -1), ""))
		cplx := minLen * codeNum
		sum += cplx
	}
	return sum
}

type pair struct {
	start string
	end   string
}

type pairDepth struct {
	pair  pair
	depth int
}

func dist(x string, y string, depth int, cache map[pairDepth]int) int {
	if depth == 1 {
		seqs := expandPair(x, y, dirPad)
		return len(seqs[0])
	}

	minLen := 9999999999999 // has to be big enough such that computed values guaranteed to be lower
	seqs := expandPair(x, y, dirPad)
	for _, seq := range seqs {
		l := 0
		seq = append([]string{"A"}, seq...)
		for i := 0; i < len(seq)-1; i++ {
			a := seq[i]
			b := seq[i+1]
			p := pairDepth{
				pair:  pair{a, b},
				depth: depth - 1,
			}
			v, ok := cache[p]
			if ok {
				l += v
			} else {
				add := dist(a, b, depth-1, cache)
				cache[p] = add
				l += add
			}
		}
		minLen = min(minLen, l)
	}
	return minLen
}

func expandPair(x string, y string, keys map[string]pos) [][]string {
	var moves [][][]string
	var allMoves [][]string
	cur := keys[x]
	next := keys[y]
	var steps []string
	diff := pos{next.i - cur.i, next.j - cur.j}

	if next == cur {
		steps = append(steps, "A")
	} else {
		var move string
		if diff.i < 0 {
			move = "^"
		} else if diff.i > 0 {
			move = "v"
		}
		for i := 0; i < absInt(diff.i); i++ {
			steps = append(steps, move)
		}

		if diff.j < 0 {
			move = "<"
		} else if diff.j > 0 {
			move = ">"
		}
		for j := 0; j < absInt(diff.j); j++ {
			steps = append(steps, move)
		}
	}

	stepPerms := permutations(steps)

	var validMoves [][]string
	for _, perm := range stepPerms {
		ok := true
		before := cur
		for _, p := range perm {
			d := disp[p]
			after := pos{before.i + d.i, before.j + d.j}
			if after == keys["X"] || !keyExists(after, keys) {
				ok = false
				break
			}
			before = after
		}

		if ok {
			if len(perm) == 1 && perm[0] == "A" {
				validMoves = append(validMoves, perm)
			} else {
				perm = append(perm, "A")
				validMoves = append(validMoves, perm)
			}
		}
	}
	moves = append(moves, validMoves)
	allMoves = cartesian(moves...)
	return allMoves
}
