package main

import (
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

// todo: fig out solution that scan/s iterates through bytes instead of reading all into mem

func solvePart1(input string) int {
	mp := parseMap(input)
	checksum := calcMovedChecksum(mp)
	return checksum
}

func parseMap(input string) []string {
	mpRaw, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	mpStr := strings.Split(string(mpRaw), "")

	var mp []string
	for i, n := range mpStr {
		blockLen, _ := strconv.Atoi(n)
		if i%2 == 0 {
			id := strconv.Itoa(i / 2)
			k := 0
			for {
				if k == blockLen {
					break
				}
				mp = append(mp, id)
				k++
			}
		} else {
			k := 0
			for {
				if k == blockLen {
					break
				}
				mp = append(mp, ".")
				k++
			}
		}
	}
	return mp
}

func calcMovedChecksum(mp []string) int {
	l := 0
	for i := l; i < len(mp); i++ {
		if mp[i] == "." {
			l = i
			break
		}
	}

	r := len(mp) - 1
	for j := r; j >= 0; j-- {
		if mp[j] != "." {
			r = j
			break
		}
	}

	for {
		if l >= r {
			break
		}

		if mp[l] == "." && mp[r] != "." {
			mp[l] = mp[r]
			mp[r] = "."
			l++
			r--
		} else if mp[l] != "." {
			l++
		} else if mp[r] == "." {
			r--
		}
	}

	checksum := 0
	for pos, idStr := range mp {
		if idStr == "." {
			break
		}
		id, _ := strconv.Atoi(idStr)
		checksum += pos * id
	}
	return checksum
}

func solvePart2(input string) int {
	mp := parseMap(input)
	checksum := calcChunkedChecksum(mp)
	return checksum
}

func calcChunkedChecksum(mp []string) int {
	l := 0
	for {
		if mp[l] == "." {
			break
		}
		l++
	}

	r := len(mp) - 1
	for {
		if mp[r] != "." {
			break
		}
		r--
	}

	locs := make(map[int][]int)
	for i, file := range mp {
		if file != "." {
			fileInt, _ := strconv.Atoi(file)
			locs[fileInt] = append(locs[fileInt], i)
		}
	}

	f, _ := strconv.Atoi(mp[r])
	for file := f; file > 0; file-- {
		fileStart := locs[file][0]
		l = 0
		for {
			for {
				if mp[l] == "." {
					break
				}
				l++
			}

			if l >= fileStart {
				break
			}

			freeLen := 0
			for {
				if mp[l+freeLen] != "." {
					break
				}
				freeLen++
			}

			spots := locs[file]
			fileLen := len(spots)

			if freeLen >= fileLen {
				for _, spot := range spots {
					mp[l] = mp[spot]
					mp[spot] = "."
					l++
				}
				break
			} else {
				for {
					if mp[l] != "." {
						break
					}
					l++
				}
				for {
					if mp[l] == "." {
						break
					}
					l++
				}
			}
		}
	}

	checksum := 0
	for pos, idStr := range mp {
		if idStr != "." {
			id, _ := strconv.Atoi(idStr)
			checksum += pos * id
		}
	}
	return checksum
}
