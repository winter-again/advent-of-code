package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
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

		ans = solvePart1("./input_smpl_2.txt")
		fmt.Println("Answer (sample_2):", ans)

		ans = solvePart1("./input.txt")
		fmt.Println("Answer:", ans)
	} else {
		ans := solvePart2("./input.txt")
		fmt.Println("Answer:", ans)
	}
}

func solvePart1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wires, gates := parseGates(file)
	eval(wires, gates)
	_, ans := parseWires(wires, "z")
	return ans
}

func parseWires(w map[string]int, p string) (string, int) {
	var wires []string
	for wire := range w {
		if strings.HasPrefix(wire, p) {
			wires = append(wires, wire)
		}
	}
	sort.Slice(wires, func(i int, j int) bool {
		return wires[i] > wires[j]
	})

	var sb strings.Builder
	for i := range wires {
		valStr := strconv.Itoa(w[wires[i]])
		sb.WriteString(valStr)
	}
	dec, _ := strconv.ParseInt(sb.String(), 2, 64)
	return sb.String(), int(dec)
}

func eval(wires map[string]int, gates []gate) {
	queue := make([]gate, len(gates))
	copy(queue, gates)

	for len(queue) > 0 {
		gate := queue[0]
		queue = queue[1:]
		_, inp1Ready := wires[gate.inp1]
		_, inp2Ready := wires[gate.inp2]

		if inp1Ready && inp2Ready {
			var res int
			if gate.op == "AND" {
				res = wires[gate.inp1] & wires[gate.inp2]
			} else if gate.op == "OR" {
				res = wires[gate.inp1] | wires[gate.inp2]
			} else {
				res = wires[gate.inp1] ^ wires[gate.inp2]
			}
			wires[gate.out] = res
		} else {
			queue = append(queue, gate)
		}
	}
}

type gate struct {
	inp1 string
	inp2 string
	op   string
	out  string
}

func parseGates(file *os.File) (map[string]int, []gate) {
	scanner := bufio.NewScanner(file)
	scanner.Split(splitGates)

	wires := make(map[string]int)
	scanner.Scan()
	inits := strings.Split(strings.TrimSpace(scanner.Text()), "\n")
	for _, init := range inits {
		fields := strings.Split(init, ": ")
		wire := fields[0]
		val, _ := strconv.Atoi(fields[1])
		wires[wire] = val
	}

	scanner.Scan()
	var gates []gate
	lines := strings.Split(strings.TrimSpace(scanner.Text()), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		gate := gate{
			inp1: fields[0],
			inp2: fields[2],
			out:  fields[4],
			op:   fields[1],
		}
		gates = append(gates, gate)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return wires, gates
}

func splitGates(data []byte, atEOF bool) (advance int, token []byte, err error) {
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

func solvePart2(input string) string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, gates := parseGates(file)
	// log.Println("WIRES", wires)

	gmap := make(map[string]gate, len(gates))
	for _, gate := range gates {
		gmap[gate.out] = gate
	}

	swaps := make([]string, 8)

	// NOTE: approach from hyperneutrino
	// NOTE: FAIL 1) line 141: kgd OR kqf -> z10
	// z10 should be getting output of an XOR op instead (z10 = carry10 XOR (x10 XOR y10))
	// line 211: y10 XOR x10 -> sqr = interm XOR for bit 10 (z10 = carry10 XOR sqr)
	// line 147: sqr XOR msq -> mwk
	// thus, msq is carry10 and this gate should point to z10 instead of mwk
	// SWAP 1) z10 <-> mwk
	z10 := gmap["z10"]
	mwk := gmap["mwk"]
	gmap["z10"] = gate{inp1: mwk.inp1, inp2: mwk.inp2, op: mwk.op, out: "z10"}
	gmap["mwk"] = gate{inp1: z10.inp1, inp2: z10.inp2, op: z10.op, out: "mwk"}
	swaps[0] = "z10"
	swaps[1] = "mwk"

	// NOTE: FAIL 2) line 129: y18 AND x18 -> z18
	// z18 should be getting res of an XOR (z18 = carry18 XOR (x18 XOR y18))
	// line 221: x18 XOR y18 -> nqq (z18 = carry18 XOR nqq)
	// line 179: nfh XOR nqq -> qgd (z18 = nfh XOR nqq)
	// thus, nfh is carry18 and this gate should point to z18 instead of qgd
	// SWAP 2) z18 <-> qgd
	z18 := gmap["z18"]
	qgd := gmap["qgd"]
	gmap["z18"] = gate{inp1: qgd.inp1, inp2: qgd.inp2, op: qgd.op, out: "z18"}
	gmap["qgd"] = gate{inp1: z18.inp1, inp2: z18.inp2, op: z18.op, out: "qgd"}
	swaps[2] = "z18"
	swaps[3] = "qgd"

	// NOTE: FAIL 3) either jmh or mrs is supp to be interm XOR
	// line 175: jmh XOR mrs -> z24 --> one must be carry and the other is interm XOR
	// only other op is bmv OR hsw -> mtn, which implies mtn is carry calc
	// and hsw is either direct carry or recarry result
	// line 201: jmh AND mrs -> bmv
	// line 283: mtn XOR tff -> z25
	// line 213: sdd OR hcp -> mrs while sdd = y23 AND x23 and hcp = gwq AND wrn
	// this is clearly a carry calc
	// so mrs could be the carry and jmh is interm XOR (x24 XOR y24)?
	// see line 156: y24 XOR x24 -> hsw; this is the intermediate XOR operation
	// SWAP 3) jmh <-> hsw
	jmh := gmap["jmh"]
	hsw := gmap["hsw"]
	gmap["jmh"] = gate{inp1: hsw.inp1, inp2: hsw.inp2, op: hsw.op, out: "jmh"}
	gmap["hsw"] = gate{inp1: jmh.inp1, inp2: jmh.inp2, op: jmh.op, out: "hsw"}
	swaps[4] = "jmh"
	swaps[5] = "hsw"

	// NOTE: FAIL 4) line 272: cvt AND wwp -> z33
	// z33 should be getting output of an XOR op (z33 = carry33 XOR (x33 XOR y33))
	// line 196: y33 XOR x33 -> wwp (z33 = carry33 XOR wwp)
	// line 302: wwp XOR cvt -> gqp (z33 = cvt XOR wwp)
	// SWAP 4) z33 <-> gqp
	z33 := gmap["z33"]
	gqp := gmap["gqp"]
	gmap["z33"] = gate{inp1: gqp.inp1, inp2: gqp.inp2, op: gqp.op, out: "z33"}
	gmap["gqp"] = gate{inp1: z33.inp1, inp2: z33.inp2, op: z33.op, out: "gqp"}
	swaps[6] = "z33"
	swaps[7] = "gqp"

	for i := 0; i < 46; i++ {
		ok := checkZWire(fmt.Sprintf("z%02d", i), i, gmap)
		if !ok {
			break
		}
		log.Printf("-------- z wire z%02d OK --------", i)
	}
	slices.Sort(swaps)
	return strings.Join(swaps, ",")
}

func checkZWire(w string, n int, gmap map[string]gate) bool {
	log.Printf("check z wire: %s - %d", w, n)
	gate := gmap[w]
	if gate.op != "XOR" {
		log.Printf("Op is %s instead of %s", gate.op, "XOR")
		return false
	}
	if n == 0 {
		return (gate.inp1 == "x00" && gate.inp2 == "y00") || (gate.inp1 == "y00" && gate.inp2 == "x00")
	} else {
		return (checkIntermXOR(gate.inp1, n, gmap) && checkCarry(gate.inp2, n, gmap)) || (checkIntermXOR(gate.inp2, n, gmap) && checkCarry(gate.inp1, n, gmap))
	}
}

func checkIntermXOR(w string, n int, gmap map[string]gate) bool {
	log.Printf("check interm XOR: %s - %d", w, n)
	gate := gmap[w]
	if gate.op != "XOR" {
		log.Printf("Op is %s instead of %s", gate.op, "XOR")
		return false
	}
	return (gate.inp1 == fmt.Sprintf("x%02d", n) && gate.inp2 == fmt.Sprintf("y%02d", n)) || (gate.inp1 == fmt.Sprintf("y%02d", n) && gate.inp2 == fmt.Sprintf("x%02d", n))
}

func checkCarry(w string, n int, gmap map[string]gate) bool {
	log.Printf("check carry: %s - %d", w, n)
	gate := gmap[w]
	if n == 1 {
		if gate.op != "AND" {
			log.Printf("Op is %s instead of %s", gate.op, "AND")
			return false
		}
		return (gate.inp1 == "x00" && gate.inp2 == "y00") || (gate.inp1 == "y00" && gate.inp2 == "x00")
	}
	if gate.op != "OR" {
		log.Printf("Op is %s instead of %s", gate.op, "OR")
		return false
	}
	return (checkDirCarry(gate.inp1, n-1, gmap) && checkRecarry(gate.inp2, n-1, gmap)) || (checkDirCarry(gate.inp2, n-1, gmap) && checkRecarry(gate.inp1, n-1, gmap))
}

func checkDirCarry(w string, n int, gmap map[string]gate) bool {
	log.Printf("check dir carry: %s - %d", w, n)
	gate := gmap[w]
	if gate.op != "AND" {
		log.Printf("Op is %s instead of %s", gate.op, "AND")
		return false
	}
	return (gate.inp1 == fmt.Sprintf("x%02d", n) && gate.inp2 == fmt.Sprintf("y%02d", n)) || (gate.inp1 == fmt.Sprintf("y%02d", n) && gate.inp2 == fmt.Sprintf("x%02d", n))
}

func checkRecarry(w string, n int, gmap map[string]gate) bool {
	log.Printf("check recarry: %s - %d", w, n)
	gate := gmap[w]
	if gate.op != "AND" {
		log.Printf("Op is %s instead of %s", gate.op, "AND")
		return false
	}
	return (checkIntermXOR(gate.inp1, n, gmap) && checkCarry(gate.inp2, n, gmap)) || (checkIntermXOR(gate.inp2, n, gmap) && checkCarry(gate.inp1, n, gmap))
}
